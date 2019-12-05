package controllers

import (
	"api.sanghoffice/components"
	"github.com/astaxie/beego/context"
)

const (
	HEADER_MESSAGE   = "Message"
	HEADER_SESSIONID = "SessionID"
)

const (
	STATUSCODE_SUCCESS, MESSAGE_SUCCESS                   = 200, "Success."
	STATUSCODE_INVALIDREQUEST, MESSAGE_INVALIDREQUEST     = 400, "Invalid request: "
	STATUSCODE_INVALIDSESSIONID, MESSAGE_INVALIDSESSIONID = 401, "Invalid session id."
	STATUSCODE_CONFLICT, MESSAGE_CONFLICT                 = 490, "Invalid Operation: "
	STATUSCODE_EXCEPTIONOCCUR, MESSAGE_EXCEPTIONOCCUR     = 500, "Exception occur: "
)

var TestingKey = ""

type Ctrl interface {
	Context() *context.Context
	PtrData() *map[interface{}]interface{}
	ServeJson()
}

func ReplyError(ctrl interface{ Ctrl }, code int, message string) {
	ctrl.Context().Output.SetStatus(code)
	ctrl.Context().Output.Header(HEADER_MESSAGE, message)
	ctrl.ServeJson()
}

func ReplySuccess(ctrl interface{ Ctrl }, json interface{}) {
	(*ctrl.PtrData())["json"] = json
	ctrl.Context().Output.Header("Message", "Success")
	ctrl.ServeJson()
}

// func CheckSessionID(ctrl interface{ Ctrl }) (string, bool) {
// 	sessionID := ctrl.Context().Input.Header(HEADER_SESSIONID)
// 	if false == components.IsValidSession(sessionID) {
// 		ctrl.Context().Output.SetStatus(STATUSCODE_INVALIDSESSIONID)
// 		ctrl.Context().Output.Header(HEADER_MESSAGE, MESSAGE_INVALIDSESSIONID)
// 		ctrl.ServeJson()
// 		return "", false
// 	}
// 	return sessionID, true
// }

// func GetUserID(ctrl interface{ Ctrl }) (string, bool) {
// 	sessionID := ctrl.Context().Input.Header(HEADER_SESSIONID)
// 	return components.GetUserID(sessionID)
// }

func FilterSessionID(ctx *context.Context) {
	sessionID := ctx.Input.Header(HEADER_SESSIONID)
	if false == components.IsValidSession(sessionID) {
		ctx.Output.SetStatus(STATUSCODE_INVALIDSESSIONID)
		ctx.Output.Header(HEADER_MESSAGE, MESSAGE_INVALIDSESSIONID)
	}
	return
}

