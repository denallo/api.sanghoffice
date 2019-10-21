package controllers

import (
	"github.com/astaxie/beego"
)

type ResidentController struct {
	beego.Controller
}

// @router / [post]
func (ctrl *ResidentController) AddResident() {
}
