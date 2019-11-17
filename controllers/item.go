package controllers

import (
	"fmt"

	"api.sanghoffice/models"
	"api.sanghoffice/tools"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
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
