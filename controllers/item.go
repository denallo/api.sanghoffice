package controllers

import (
	"fmt"

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
	year, _ := this.GetInt("year")
	month, _ := this.GetInt("month")
	brief, success := models.GetBrief(year, month)
	if !success {
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR, MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf(""))
		return
	}
	json := map[string]interface{}{}
	json["brief"] = brief
	ReplySuccess(this, json)
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
				"WHERE  resident_id = ? WHERE cancel = 0")
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
