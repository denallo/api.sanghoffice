package controllers

import (
	"fmt"
	"strconv"

	"api.sanghoffice/models"
	"api.sanghoffice/tools"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/bitly/go-simplejson"
)

// Operations about object
type KutiController struct {
	beego.Controller
}

func (ctrl *KutiController) Context() *context.Context {
	return ctrl.Ctx
}

func (ctrl *KutiController) ServeJson() {
	ctrl.ServeJSON()
}

func (ctrl *KutiController) PtrData() *map[interface{}]interface{} {
	return &(ctrl.Data)
}

// @router / [get]
func (ctrl *KutiController) Get() {
	kutiForSex, _ := strconv.Atoi(ctrl.Ctx.Input.Param("sex"))
	retJson := models.GetKuties(kutiForSex)
	ctrl.Data["json"] = retJson
	ctrl.ServeJSON()
}

// @router /status [patch]
func (this *KutiController) UpdateBrokenStatus() {
	js, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
	jsMap, _ := js.Map()
	kutiNumber, _ := tools.JsonNumberToInt(jsMap["kutiNumber"])
	kutiType, _ := tools.JsonNumberToInt(jsMap["kutiType"])
	forSex, _ := tools.JsonNumberToInt(jsMap["forSex"])
	brokenStatus, _ := tools.JsonNumberToInt(jsMap["brokenStatus"])
	if true == models.UpdateBrokenStatus(kutiNumber, kutiType, forSex, brokenStatus) {
		ReplySuccess(this, nil)
	} else {
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR, MESSAGE_EXCEPTIONOCCUR+fmt.Sprintf("更新Broken状态失败"))
	}
}
