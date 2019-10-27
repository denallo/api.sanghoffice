package controllers

import (
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

// @router / [post]
func (this *ResidentController) AddResident() {

}
