package controllers

import (
	"api.sanghoffice/components"
	"api.sanghoffice/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
)

type UsersCtrl struct {
	beego.Controller
}

func (ctrl *UsersCtrl) Context() *context.Context {
	return ctrl.Ctx
}

func (ctrl *UsersCtrl) ServeJson() {
	ctrl.ServeJSON()
}

func (ctrl *UsersCtrl) PtrData() *map[interface{}]interface{} {
	return &(ctrl.Data)
}

//@router / [post]
func (this *UsersCtrl) Login() {
	js, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
	jsMap, _ := js.Map()
	username, _ := jsMap["UserName"].(string)
	password, _ := jsMap["Password"].(string)
	fingerprint, _ := jsMap["Fingerprint"].(string)
	o := orm.NewOrm()
	user := models.User{
		UserName: username,
		Password: password}
	err := o.Read(&user, "username", "password")
	if nil != err {
		println(err.Error())
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+err.Error())
		return
	}
	sessionID := components.GenerateSessionID(fingerprint, false, "")
	retJson := map[string]string{}
	retJson["sessionID"] = sessionID
	ReplySuccess(this, retJson)
}

//@router / [get]
func (this *UsersCtrl) QuerySessionID() {
	fingerprint := this.GetString("Fingerprint")
	sessionID, existed := components.GetSessionID(fingerprint)
	if !existed {
		ReplyError(this, 404, "session id not found")
		return
	}
	retJson := map[string]string{}
	retJson["sessionID"] = sessionID
	ReplySuccess(this, retJson)
}
