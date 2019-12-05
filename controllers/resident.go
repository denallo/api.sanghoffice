package controllers

import (
	"fmt"

	"api.sanghoffice/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/bitly/go-simplejson"
)

type ResidentCtrl struct {
	beego.Controller
}

func (this *ResidentCtrl) Context() *context.Context {
	return this.Ctx
}

func (this *ResidentCtrl) ServeJson() {
	this.ServeJSON()
}

func (this *ResidentCtrl) PtrData() *map[interface{}]interface{} {
	return &(this.Data)
}

// @router / [get]
func (this *ResidentCtrl) GetResidents() {
	sex, _ := this.GetInt("sex")
	userRole := UserRole(this)
	switch userRole {
	case models.ROLE_MALE_ADM:
	case models.ROLE_FEMALE_ADM:
		sex = userRole
	}
	state, _ := this.GetInt("state")
	residents, success := models.GetResidents(sex, state)
	if !success {
		ReplyError(this,
			STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf(""))
		return
	}
	json := map[string]interface{}{}
	json["residents"] = residents
	ReplySuccess(this, json)
}

// @router /info [get]
func (this *ResidentCtrl) GetResidentInfo() {
	id, _ := this.GetInt("id")
	info, success := models.GetResidentInfo(id)
	if !success {
		ReplyError(this,
			STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf(""))
		return
	}
	json := map[string]interface{}{}
	json["info"] = info
	ReplySuccess(this, json)
}

// @router /info [patch]
func (this *ResidentCtrl) UpdateResidentInfo() {
	js, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
	jsMap, _ := js.Map()
	info, success := models.UpdateResident(jsMap)
	if !success {
		ReplyError(this,
			STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf(""))
		return
	}
	json := map[string]interface{}{}
	json["info"] = info
	ReplySuccess(this, json)
}
