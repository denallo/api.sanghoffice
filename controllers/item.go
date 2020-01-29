package controllers

import (
	"fmt"
	"time"

	"api.sanghoffice/models"
	"api.sanghoffice/tools"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
)

type ItemController struct {
	beego.Controller
}

func (ctrl *ItemController) Context() *context.Context {
	return ctrl.Ctx
}

func (ctrl *ItemController) ServeJson() {
	ctrl.ServeJSON()
}

func (ctrl *ItemController) PtrData() *map[interface{}]interface{} {
	return &(ctrl.Data)
}

// @router / [get]
func (this *ItemController) GetBrief() {
	userRole := UserRole(this)
	sex := -1
	switch userRole {
	case models.ROLE_MALE_ADM:
	case models.ROLE_FEMALE_ADM:
		sex = userRole
	}
	_, brief := this.GetInt("brief")
	_, unconfirmed := this.GetInt("unconfirmed")
	if brief == nil {
		year, _ := this.GetInt("year")
		month, _ := this.GetInt("month")
		brief, success := models.GetBrief(year, month, sex)
		if !success {
			ReplyError(this, STATUSCODE_EXCEPTIONOCCUR, MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf(""))
			return
		}
		json := map[string]interface{}{}
		json["brief"] = brief
		ReplySuccess(this, json)
	} else if unconfirmed == nil {
		residentID, _ := this.GetInt("residentID")
		o := orm.NewOrm()
		currDate := time.Now().Format("2006-01-02")
		items := []*models.Item{}
		query := o.QueryTable("tb_item").
			Filter("resident_id", residentID).
			Filter("enabled", 1).Filter("confirmed", 0).
			Filter("type__in",
				models.TYPE_APPOINT_TO_ARRIVE,
				models.TYPE_PLAN_TO_LEAVE).
			Filter("activate_date__lte", currDate)
		_, err := query.All(&items)
		if nil != err {
			println(err.Error())
			ReplyError(this,
				STATUSCODE_EXCEPTIONOCCUR,
				MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf("%s", err.Error()))
			return
		}
		retJson := map[string]interface{}{}
		retJson["residentID"] = residentID
		retItems := []map[string]interface{}{}
		for i := 0; i < len(items); i++ {
			retItem := map[string]interface{}{}
			retItem["itemID"] = items[i].Id
			retItem["itemType"] = items[i].Type
			retItem["activateDate"] = items[i].ActivateDate
			retItems = append(retItems, retItem)
		}
		retJson["unconfirmeds"] = retItems
		ReplySuccess(this, retJson)
	}
}

// @router /actions/confirm [patch]
func (this *ItemController) Confirm() {
	js, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
	jsMap, _ := js.Map()
	residentID, _ := tools.JsonNumberToInt(jsMap["residentID"])
	stateType, _ := tools.JsonNumberToInt(jsMap["stateType"])
	success := models.UpdateResidentState(residentID, stateType)
	if !success {
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf(""))
		return
	}
	ReplySuccess(this, nil)
}

// @router /residents/actions/leave [patch]
func (this *ItemController) Leave() {
	js, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
	jsMap, _ := js.Map()
	residentID, _ := tools.JsonNumberToInt(jsMap["residentID"])
	var err error
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		println(err.Error())
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+fmt.Sprint(err.Error()))
		return
	}
	for i := 0; i < 1; i++ { // golang没有try-catch，这里用for来凑合下
		if nil != err {
			println(err.Error())
			break
		}
		// 所有相关item置为enabled和confirmed
		sql := fmt.Sprintf(
			"UPDATE tb_item SET enabled = 1, confirmed = 1 " +
				"WHERE  resident_id = ? AND canceled = 0")
		_, err = o.Raw(sql, residentID).Exec()
		if nil != err {
			break
		}
		// 解除住众与孤邸的绑定（删除tb_resi_status）
		_, err = o.Delete(&models.ResiStatus{ResidentId: residentID})
		if nil != err {
			break
		}
		break
	}
	if err != nil {
		println(err.Error())
		o.Rollback()
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+fmt.Sprint(err.Error()))
	} else {
		o.Commit()
		ReplySuccess(this, "")
	}

}

// @router /appointments/actions/cancel [patch]
func (this *ItemController) CancelAppointment() {
	js, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
	jsMap, _ := js.Map()
	residentID, _ := tools.JsonNumberToInt(jsMap["residentID"])
	println(residentID)
	var err error
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		println(err.Error())
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+fmt.Sprint(err.Error()))
		return
	}
	for i := 0; i < 1; i++ { // golang没有try-catch，这里用for来凑合下
		if nil != err {
			println(err.Error())
			break
		}
		// 所有相关item置为canceled
		sql := fmt.Sprintf(
			"UPDATE tb_item SET canceled = 1 " +
				"WHERE  resident_id = ? AND confirmed = 0")
		_, err = o.Raw(sql, residentID).Exec()
		if nil != err {
			break
		}
		// 解除住众与孤邸的绑定（删除tb_resi_status）
		_, err = o.Delete(&models.ResiStatus{ResidentId: residentID})
		if nil != err {
			break
		}
		break
	}
	if err != nil {
		println(err.Error())
		o.Rollback()
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+fmt.Sprint(err.Error()))
	} else {
		o.Commit()
		ReplySuccess(this, "")
	}
}

// @router /residents/actions/change-leaving [patch]
func (this *ItemController) ChangeLeavingDate() {
	js, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
	jsMap, _ := js.Map()
	residentID, _ := tools.JsonNumberToInt(jsMap["residentID"])
	date := jsMap["date"].(string)
	println(residentID, date)
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		println(err.Error())
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+fmt.Sprint(err.Error()))
		return
	}
	for i := 0; i < 1; i++ { // golang没有try-catch，这里用for来凑合下
		if nil != err {
			println(err.Error())
			break
		}
		// 更新相关item的activate_date
		sql := fmt.Sprintf(
			"UPDATE tb_item SET activate_date = ? " +
				"WHERE  resident_id = ? AND canceled = 0 AND " +
				"(enabled=1 AND confirmed=0) AND type=?")
		_, err = o.Raw(sql,
			date, residentID,
			models.TYPE_PLAN_TO_LEAVE).Exec()
		if nil != err {
			println(err.Error())
			break
		}
		// 修改resident的离寺日期字段
		sql = fmt.Sprintf(
			"UPDATE tb_resi_status SET plan_to_leave_date = ? " +
				"WHERE resident_id = ?")
		_, err = o.Raw(sql, date, residentID).Exec()
		if nil != err {
			println(err.Error())
			break
		}
		break
	}
	if err != nil {
		println(err.Error())
		o.Rollback()
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+fmt.Sprint(err.Error()))
	} else {
		o.Commit()
		ReplySuccess(this, "")
	}
}

// @router /appointments/actions/change [patch]
func (this *ItemController) ChangeAppointedDate() {
	js, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
	jsMap, _ := js.Map()
	residentID, _ := tools.JsonNumberToInt(jsMap["residentID"])
	date := jsMap["date"].(string)
	println(residentID, date)
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		println(err.Error())
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+fmt.Sprint(err.Error()))
		return
	}
	for i := 0; i < 1; i++ { // golang没有try-catch，这里用for来凑合下
		if nil != err {
			println(err.Error())
			break
		}
		// 更新相关item的activate_date
		sql := fmt.Sprintf(
			"UPDATE tb_item SET activate_date = ? " +
				"WHERE  resident_id = ? AND canceled = 0 AND " +
				"(enabled=1 AND confirmed=0) AND type=?")
		_, err = o.Raw(sql,
			date, residentID,
			models.TYPE_APPOINT_TO_ARRIVE).Exec()
		if nil != err {
			println(err.Error())
			break
		}
		// 修改resident的到达日期字段
		sql = fmt.Sprintf(
			"UPDATE tb_resi_status SET arrive_date = ? " +
				"WHERE resident_id = ?")
		_, err = o.Raw(sql, date, residentID).Exec()
		if nil != err {
			println(err.Error())
			break
		}
		break
	}
	if err != nil {
		println(err.Error())
		o.Rollback()
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+fmt.Sprint(err.Error()))
	} else {
		o.Commit()
		ReplySuccess(this, "")
	}
}
