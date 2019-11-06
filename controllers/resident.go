package controllers

import (
	"fmt"

	"api.sanghoffice/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type ResidentController struct {
	beego.Controller
}

func (this *ResidentController) Context() *context.Context {
	return this.Ctx
}

func (this *ResidentController) ServeJson() {
	this.ServeJSON()
}

func (this *ResidentController) PtrData() *map[interface{}]interface{} {
	return &(this.Data)
}

// @router / [get]
func (this *ResidentController) GetResidents() {
	sex, _ := this.GetInt("sex")
	residents, success := models.GetResidents(sex)
	if !success {
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR, MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf(""))
	}
	json := map[string]interface{}{}
	json["residents"] = residents
	ReplySuccess(this, json)
}
