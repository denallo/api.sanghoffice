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
	// 根据用户角色进行相应处理
	role := user.Role
	success, info := components.SetUserData(sessionID, components.USERDATA_ROLE, role)
	if !success {
		println(info)
		ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
			MESSAGE_EXCEPTIONOCCUR+info)
		return
	}
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

//@router / [delete]
func (this *UsersCtrl) Logout() {
	// js, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
	// jsMap, _ := js.Map()
	// fingerprint, _ := jsMap["Fingerprint"].(string)
	fingerprint := this.GetString("Fingerprint")
	if components.IsRegistedUser(fingerprint) {
		components.DelSession(fingerprint)
	}
	ReplySuccess(this, nil)
	println("success")
}
