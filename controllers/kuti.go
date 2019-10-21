package controllers

import (
	"strconv"

	"api.sanghoffice/models"
	"github.com/astaxie/beego"
)

// Operations about object
type KutiController struct {
	beego.Controller
}

// @router / [get]
func (ctrl *KutiController) Get() {
	kutiForSex, _ := strconv.Atoi(ctrl.Ctx.Input.Param("sex"))
	retJson := models.GetKuties(kutiForSex)
	ctrl.Data["json"] = retJson
	ctrl.ServeJSON()
}
