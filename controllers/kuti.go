package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"

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
func (this *KutiController) Get() {
	kutiForSex, _ := strconv.Atoi(this.Ctx.Input.Param("sex"))
	_, getRange := this.GetInt("range")
	_, getValidTypes := this.GetInt("types")
	retJson := map[string]interface{}{}
	o := orm.NewOrm()
	if nil == getRange {
		kutiType, _ := this.GetInt("type")
		kuties := []*models.Kuti{}
		cnt, err := o.QueryTable("tb_kuti").
			Filter("for_sex", kutiForSex).
			Filter("broken", 0).
			Filter("type", kutiType).
			OrderBy("number").
			All(&kuties)
		if nil != err {
			ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
				MESSAGE_EXCEPTIONOCCUR+err.Error())
			return
		} else if 0 == cnt {
			ReplyError(this,
				STATUSCODE_INVALIDREQUEST,
				MESSAGE_INVALIDREQUEST+fmt.Sprintf(
					"found no record of type %d", kutiType))
			return
		}
		eliminatedKuties := []*models.Kuti{}
		_, err = o.QueryTable("tb_kuti").
			Filter("for_sex", kutiForSex).
			Filter("broken", 1).
			Filter("type", kutiType).
			OrderBy("number").
			All(&eliminatedKuties)
		if nil != err {
			ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
				MESSAGE_EXCEPTIONOCCUR+err.Error())
			return
		}
		minNumber := kuties[0].Number
		maxNumber := kuties[len(kuties)-1].Number
		eliminated := []int{}
		for i := 0; i < len(eliminatedKuties); i++ {
			eliminated = append(eliminated, eliminatedKuties[i].Number)
		}
		retJson["from"] = minNumber
		retJson["to"] = maxNumber
		retJson["eliminated"] = eliminated
	} else if nil == getValidTypes {
		validTypes := []int{}
		_, err := o.Raw("SELECT DISTINCT type FROM tb_kuti "+
			"WHERE for_sex=? AND broken=0 ORDER BY type",
			kutiForSex).QueryRows(&validTypes)
		if nil != err {
			ReplyError(this, STATUSCODE_EXCEPTIONOCCUR,
				MESSAGE_EXCEPTIONOCCUR+err.Error())
			return
		}
		retJson["validTypes"] = validTypes
	} else {
		retJson = models.GetKuties(kutiForSex)
	}
	this.Data["json"] = retJson
	this.ServeJSON()
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
