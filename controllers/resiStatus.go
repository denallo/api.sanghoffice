package controllers

import (
	"encoding/json"
	"fmt"

	"api.sanghoffice/models"
	"github.com/astaxie/beego"
)

type ResiStatusCtrl struct {
	beego.Controller
}

// @router / [post]
func (ctrl *ResiStatusCtrl) AddResiStatus() {
	var req models.ReqNewResiStatus
	json.Unmarshal(ctrl.Ctx.Input.RequestBody, &req)
	println(fmt.Sprintf("%+v", req))
	retJson := map[string]interface{}{}
	ctrl.Data["json"] = retJson
	ctrl.ServeJSON()
}
