package controllers

import (
	"fmt"

	"api.sanghoffice/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
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
