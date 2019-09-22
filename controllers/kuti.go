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
func (o *KutiController) Get() {
	kutiForSex, _ := strconv.Atoi(o.Ctx.Input.Param("sex"))
	kuties := models.GetKuties(kutiForSex)
	o.Data["json"] = kuties
	o.ServeJSON()
}
