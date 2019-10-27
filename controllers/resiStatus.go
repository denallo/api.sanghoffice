package controllers

import (
	"fmt"

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
	var isExisted bool
	var residentID int
	sex, _ := tools.JsonNumberToInt(jsMap["sex"])
	// nameText := ""
	if name != "" {
		// nameText = name
		isExisted, residentID = models.IsExistedResident(name, false, sex)
	}
	if dhamame != "" {
		// nameText = dhamame
		isExisted, residentID = models.IsExistedResident(dhamame, true, sex)
	}
	if !isExisted {
		// ReplyError(this, STATUSCODE_CONFLICT, MESSAGE_CONFLICT+fmt.Sprintf("unknow Resident：%s，cannot create ResiStatus", nameText))
		residentID = models.AddResident(jsMap)
		if -1 == residentID {
			ReplyError(this, STATUSCODE_EXCEPTIONOCCUR, MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf("fail to create resident resource"))
			return
		}
	}
	kutiNumber, _ := tools.JsonNumberToInt(jsMap["kutiNumber"])
	kutiType, _ := tools.JsonNumberToInt(jsMap["kutiType"])
	kutiIndex, _ := tools.JsonNumberToInt(jsMap["kutiIndex"])
	isMonk, _ := tools.JsonNumberToInt(jsMap["isMonk"])
	arriveDate := jsMap["arriveDate"].(string)
	leaveDate := jsMap["leaveDate"].(string)
	success := models.AddResiStatus(residentID, sex, kutiNumber, kutiType, arriveDate, leaveDate)
	if !success {
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR, MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf(""))
	}
	retJson := map[string]interface{}{}
	retJson["residentId"] = residentID
	retJson["name"] = name
	retJson["dhamame"] = dhamame
	retJson["arriveDate"] = arriveDate
	retJson["leaveDate"] = leaveDate
	retJson["isMonk"] = isMonk
	retJson["kutiNumber"] = kutiNumber
	retJson["kutiNumber"] = kutiNumber
	retJson["kutiType"] = kutiType
	retJson["kutiIndex"] = kutiIndex
	this.Data["json"] = retJson
	ReplySuccess(this, retJson)
}
