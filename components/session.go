package components

import (
	"time"

	"github.com/astaxie/beego/cache"
	uuid "github.com/satori/go.uuid"
)

var user2session, sessionData cache.Cache

const TIMEOUT = 5 * 60 * time.Second

func init() {
	user2session, _ = cache.NewCache("memory", `{"interval":60}`)
	sessionData, _ = cache.NewCache("memory", `{"interval":60}`)
}

type SessionData struct {
	PrivateKey string
	WxOpenID   string
	WxUnionID  string
	UserData   map[string]interface{}
}

const (
	USERDATA_LASTORDERID = "LastOrderID"
	USERDATA_DONORID     = "DonorID"
	USERDATA_ROLE        = "Role"
)

// 生成用户会话ID
func GenerateSessionID(userID string, isUnionID bool, sessionKey string) (sessionID string) {
	if isRegistedUser(userID) {
		delSession(userID)
	}
	_uuid, _ := uuid.NewV1()
	sessionID = _uuid.String()
	sessData := SessionData{PrivateKey: sessionKey, WxOpenID: userID, UserData: make(map[string]interface{})}
	saveSession(sessionID, sessData)
	return
}

// 检查用户会话ID合法性
func IsValidSession(sessionID string) (isValid bool) {
	isValid = sessionData.IsExist(sessionID)
	if isValid {
		refreshSession(sessionID)
	}
	return
}

// 获取用户系统唯一标识
func GetUserID(sessionID string) (userID string, success bool) {
	success = false
	res := sessionData.Get(sessionID)
	if nil == res {
		return
	}
	sessData := res.(SessionData)
	userID = sessData.WxOpenID
	success = true
	return
}

// 获取会话ID
func GetSessionID(userID string) (sessionID string, success bool) {
	success = false
	if !isRegistedUser(userID) {
		return
	}
	res := user2session.Get(userID)
	if nil == res {
		return
	}
	sessionID = res.(string)
	success = true
	return
}

// 清空缓存
func ClearCache() {
	_ = user2session.ClearAll()
	_ = sessionData.ClearAll()
}

// 获取用户会话业务数据
func GetUserData(sessionID string, key string) (data interface{}, success bool) {
	res := sessionData.Get(sessionID)
	if nil == res {
		// success = false
		return nil, false
	}
	sessData := res.(SessionData)
	data, success = sessData.UserData[key]
	return
}

// 设置用户会话业务数据
func SetUserData(sessionID string, key string, data interface{}) (bool, string) {
	res := sessionData.Get(sessionID)
	if nil == res {
		return false, "Cannot get seesion data of id " + sessionID
	}
	sessData := res.(SessionData)
	sessData.UserData[key] = data
	err := sessionData.Put(sessionID, sessData, TIMEOUT)
	if nil != err {
		// panic(err)
		return false, err.Error()
	}
	return true, ""
}

func saveSession(sessionID string, sessData SessionData) {
	sessionData.Put(sessionID, sessData, TIMEOUT)
	user2session.Put(sessData.WxOpenID, sessionID, TIMEOUT)
	return
}

func refreshSession(sessionID string) {
	sessData := sessionData.Get(sessionID)
	if nil == sessData {
		return
	}
	saveSession(sessionID, sessData.(SessionData))
}

func isRegistedUser(userID string) (existed bool) {
	existed = user2session.IsExist(userID)
	return
}

func delSession(userID string) {
	sessionID := user2session.Get(userID).(string)
	user2session.Delete(userID)
	sessionData.Delete(sessionID)
}
