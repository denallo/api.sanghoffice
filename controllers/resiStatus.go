package controllers

import (
	"fmt"

	"github.com/astaxie/beego/orm"

	"api.sanghoffice/models"
	"api.sanghoffice/tools"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/bitly/go-simplejson"
)

type ResiStatusCtrl struct {
	beego.Controller
}

func (ctrl *ResiStatusCtrl) Context() *context.Context {
	return ctrl.Ctx
}

func (ctrl *ResiStatusCtrl) ServeJson() {
	ctrl.ServeJSON()
}

func (ctrl *ResiStatusCtrl) PtrData() *map[interface{}]interface{} {

	return &(ctrl.Data)
}

// @router / [post]
func (this *ResiStatusCtrl) AddResiStatus() {
	js, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
	jsMap, _ := js.Map()
	name := jsMap["name"].(string)
	dhamame := jsMap["dhamame"].(string)
	identifier := jsMap["identifier"].(string)
	sex, _ := tools.JsonNumberToInt(jsMap["sex"])
	age, _ := tools.JsonNumberToInt(jsMap["age"])
	residentType, _ := tools.JsonNumberToInt(jsMap["type"])
	folk := jsMap["folk"].(string)
	nativePlace := jsMap["native_place"].(string)
	ability := jsMap["ability"].(string)
	phone := jsMap["phone"].(string)
	emergencyContact := jsMap["emergency_contact"].(string)
	emergencyContactPhone := jsMap["emergency_contact_phone"].(string)
	// var isExisted bool
	// var residentID int
	// sex, _ := tools.JsonNumberToInt(jsMap["sex"])
	// // nameText := ""
	// if name != "" {
	// 	// nameText = name
	// 	isExisted, residentID = models.IsExistedResident(name, false, sex)
	// }
	// if dhamame != "" {
	// 	// nameText = dhamame
	// 	isExisted, residentID = models.IsExistedResident(dhamame, true, sex)
	// }
	// if !isExisted {
	// 	// ReplyError(this,
	// 	// 	STATUSCODE_CONFLICT,
	// 	// 	MESSAGE_CONFLICT+fmt.Sprintf(
	// 	// 		"unknow Resident：%s，cannot create ResiStatus", nameText))
	// 	residentID = models.AddResident(jsMap)
	// 	if -1 == residentID {
	// 		ReplyError(this,
	// 			STATUSCODE_EXCEPTIONOCCUR,
	// 			MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf(
	// 				"fail to create resident resource"))
	// 		return
	// 	}
	// }

	kutiNumber, _ := tools.JsonNumberToInt(jsMap["kutiNumber"])
	kutiType, _ := tools.JsonNumberToInt(jsMap["kutiType"])
	kutiIndex, _ := tools.JsonNumberToInt(jsMap["kutiIndex"])
	isMonk, _ := tools.JsonNumberToInt(jsMap["isMonk"])
	arriveDate := jsMap["arriveDate"].(string)
	leaveDate := jsMap["leaveDate"].(string)

	residentID, success := models.CheckIn(
		name, dhamame, identifier, sex,
		age, residentType, folk,
		nativePlace, ability, phone,
		emergencyContact, emergencyContactPhone,
		kutiNumber, kutiType, isMonk,
		arriveDate, leaveDate)
	if !success {
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf(""))
		return
	}
	// success := models.AddResiStatus(
	// 	residentID, sex,
	// 	kutiNumber, kutiType,
	// 	arriveDate, leaveDate)
	// if !success {
	// 	ReplyError(this,
	// 		STATUSCODE_EXCEPTIONOCCUR,
	// 		MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf(""))
	// 	return
	// }
	// success = models.CreateItem(residentID, arriveDate, leaveDate)
	// if !success {
	// 	ReplyError(
	// 		this,
	// 		STATUSCODE_EXCEPTIONOCCUR,
	// 		MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf(""))
	// 	return
	// }
	var avaliables []([]int)
	avaliables, success = models.GetAvailablesInfo(kutiNumber, kutiType, sex)
	retJson := map[string]interface{}{}
	retJson["residentId"] = residentID
	retJson["name"] = name
	retJson["dhamame"] = dhamame
	retJson["arriveDate"] = arriveDate
	retJson["leaveDate"] = leaveDate
	retJson["isMonk"] = isMonk
	// retJson["kutiNumber"] = kutiNumber
	// retJson["kutiType"] = kutiType
	retJson["kutiIndex"] = kutiIndex
	retJson["availables"] = avaliables
	this.Data["json"] = retJson
	ReplySuccess(this, retJson)
}

// @router / [patch]
func (this *ResiStatusCtrl) ChangeKuti() {
	js, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
	jsMap, _ := js.Map()
	residentID, _ := tools.JsonNumberToInt(jsMap["residentID"])
	kutiType, _ := tools.JsonNumberToInt(jsMap["kutiType"])
	kutiNumber, _ := tools.JsonNumberToInt(jsMap["kutiNumber"])
	kutiForSex, _ := tools.JsonNumberToInt(jsMap["kutiForSex"])
	o := orm.NewOrm()
	kutiID := 0
	err := o.Raw(
		"SELECT id FROM tb_kuti "+
			"WHERE type=? AND for_sex=? AND number=?",
		kutiType, kutiForSex, kutiNumber).QueryRow(&kutiID)
	if nil != err {
		println(err.Error())
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+err.Error())
		return
	}
	_, err = o.Raw(
		"UPDATE tb_resi_status SET kuti_id=? WHERE resident_id=?",
		kutiID, residentID).Exec()
	if nil != err {
		println(err.Error())
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+err.Error())
		return
	}
	ReplySuccess(this, nil)
}
