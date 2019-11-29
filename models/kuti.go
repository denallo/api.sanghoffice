package models

import (
	"fmt"
	"strconv"
	"time"

	mapset "github.com/deckarep/golang-set"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"
const DATE_LAYOUT = "2006-01-02"

func calcAvaliables(startDate string, endDate string, arriveDate string, leaveDate string) []int {
	if arriveDate == "" {
		arriveDate = "2019-09-01 00:00:00"
	}
	if leaveDate == "" {
		leaveDate = "2100-01-01 00:00:00"
	}
	if len(arriveDate) == 10 {
		arriveDate += " 00:00:00"
	}
	if len(leaveDate) == 10 {
		leaveDate += " 00:00:00"
	}
	var avaliable []int
	dateSession := mapset.NewSet()
	dateSessionArray := []string{}
	for currIndex := startDate; currIndex <= endDate; {
		dateSessionArray = append(dateSessionArray, currIndex)
		dateSession.Add(currIndex)
		t, _ := time.Parse(TIME_LAYOUT, currIndex)
		after, _ := time.ParseDuration("24h")
		currIndex = t.Add(after).Format(TIME_LAYOUT)
	}
	engagedSession := mapset.NewSet()
	for currIndex := arriveDate; currIndex < leaveDate; {
		if currIndex < startDate {
			t, _ := time.Parse(TIME_LAYOUT, currIndex)
			after, _ := time.ParseDuration("24h")
			currIndex = t.Add(after).Format(TIME_LAYOUT)
			continue
		} else if currIndex > endDate {
			break
		}
		engagedSession.Add(currIndex)
		t, _ := time.Parse(TIME_LAYOUT, currIndex)
		after, _ := time.ParseDuration("24h")
		currIndex = t.Add(after).Format(TIME_LAYOUT)
	}
	avaliableDates := dateSession.Difference(engagedSession)
	for i := 0; i < len(dateSessionArray); i++ {
		var status int // 0-空闲 1-有人
		if avaliableDates.Contains(dateSessionArray[i]) {
			status = 0
		} else {
			status = 1
		}
		avaliable = append(avaliable, status)
	}
	return avaliable
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func getFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return getZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func getLastDateOfMonth(d time.Time) time.Time {
	return getFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//获取某一天的0点时间
func getZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func getDateSessions() []([2]string) {
	var dateSessions []([2]string)
	currTime := time.Now()
	for i := 0; i < 5; i++ {
		layout := "2006-01-02 15:04:05"
		firstDayOfMonth := getFirstDateOfMonth(currTime).Format(layout)
		lastDayOfMonth := getLastDateOfMonth(currTime).Format(layout)
		dateSession := [2]string{firstDayOfMonth, lastDayOfMonth}
		dateSessions = append(dateSessions, dateSession)
		currTime = getLastDateOfMonth(currTime).AddDate(0, 0, 1)
	}
	return dateSessions
}

func GetKuties(forSex int) map[string]interface{} {
	retJson := map[string]interface{}{}
	// 孤邸属性
	mapKuties := map[int]*Kuti{}
	var kuties []*Kuti // 按类型、孤邸号排序
	query := o.QueryTable("tb_kuti").Filter("for_sex", forSex).OrderBy("type", "number")
	num, err := query.All(&kuties)
	if nil != err {
		panic(err)
	}
	typeLeader := []int{} // kuties数组中每种类型第一个孤邸的索引
	groupCount := []int{} // 每种类型的孤邸数量
	lastType := -1
	cnt := 0
	for i := 0; i < int(num); i++ {
		if kuties[i].Type != lastType {
			if lastType != -1 {
				groupCount = append(groupCount, cnt)
				cnt = 0
			}
			typeLeader = append(typeLeader, i)
			lastType = kuties[i].Type
		}
		cnt += 1
		index := kuties[i].Id
		mapKuties[index] = kuties[i]
	}
	groupCount = append(groupCount, cnt)
	retJson["typeLeaderIndex"] = typeLeader
	retJson["typeGroupCnt"] = groupCount
	// 人员信息
	mapResidents := map[int]*Resident{}
	var residents []*Resident
	query = o.QueryTable("tb_resident").Filter("sex", forSex)
	num, err = query.All(&residents)
	if nil != err {
		panic(err)
	}
	for i := 0; i < int(num); i++ {
		index := residents[i].Id
		mapResidents[index] = residents[i]
	}
	// 入住安排情况
	var resiStatusList []*ResiStatus
	mapKuti2Status := map[int]([]*ResiStatus){}
	query = o.QueryTable("tb_resi_status")
	num, err = query.All(&resiStatusList)
	if nil != err {
		panic(err)
	}
	for i := 0; i < int(num); i++ {
		resiStatus := resiStatusList[i]
		if _, existed := mapKuties[resiStatus.KutiId]; !existed {
			continue
		}
		mapKuti2Status[resiStatus.KutiId] = append(mapKuti2Status[resiStatus.KutiId], resiStatus)
	}
	// 打包json
	var kutiesInfo [](map[string]interface{})
	for i := 0; i < len(kuties); i++ { // 遍历孤邸
		kutiInfo := kuties[i]
		unconfirmeds := []int{}
		jsonKutiInfo := map[string]interface{}{}
		jsonKutiInfo["kutiNumber"] = kutiInfo.Number
		jsonKutiInfo["type"] = kutiInfo.Type
		jsonKutiInfo["forSex"] = kutiInfo.ForSex
		jsonKutiInfo["broken"] = kutiInfo.Broken
		var listEnagedStatus []([2]string) // [(arriveDate, leaveDate), ...]
		var listResidentsInfo [](map[string]interface{})
		for j := 0; j < len(mapKuti2Status[kutiInfo.Id]); j++ { // 某孤邸入住的所有人员
			residentInfo := map[string]interface{}{}
			resiStatus := mapKuti2Status[kutiInfo.Id][j]
			resident, existed := mapResidents[resiStatus.ResidentId]
			if !existed {
				// fmt.Println(resiStatus.ResidentId)
				continue
			}
			// 入住与离开日期
			enagedStatus := [2]string{resiStatus.ArriveDate, resiStatus.PlanToLeaveDate}
			listEnagedStatus = append(listEnagedStatus, enagedStatus)
			if remainUnconfirmed(resident.Id) {
				jsonKutiInfo["blocked"] = 1
				unconfirmeds = append(unconfirmeds, resident.Id)
			}
			// 人员信息
			residentInfo["id"] = resident.Id
			residentInfo["dhamame"] = resident.Dhamame
			residentInfo["name"] = resident.Name
			if R_TYPE_BHIKHU == resident.Type ||
				R_TYPE_SAMANERA == resident.Type ||
				R_TYPE_SAYALAY == resident.Type ||
				R_TYPE_OTHER_MONK == resident.Type {
				// if resident.Dhamame != "" {
				// 	residentInfo["name"] = resident.Dhamame
				// } else {
				// 	residentInfo["name"] = resident.Name
				// }
				residentInfo["isMonk"] = 1
			} else {
				// residentInfo["name"] = resident.Name
				residentInfo["isMonk"] = 0
			}
			residentInfo["leaveDate"] = resiStatus.PlanToLeaveDate
			if residentInfo["leaveDate"] == "" {
				residentInfo["leaveDate"] = "常住人员"
			}
			residentInfo["arriveDate"] = resiStatus.ArriveDate
			listResidentsInfo = append(listResidentsInfo, residentInfo)
		}
		// 计算孤邸在某段时间内的入住安排详情
		dateSessions := getDateSessions()
		var avaliables []([]int)
		for index := 0; index < len(dateSessions); index++ {
			var avaliableUnmerged []([]int) // 用于合并一栋孤邸多个住众在同一天的入住状态
			var avaliableMerged []int
			startDate := dateSessions[index][0]
			endDate := dateSessions[index][1]
			for sub_index := 0; sub_index < len(listEnagedStatus); sub_index++ {
				arriveDate := listEnagedStatus[sub_index][0]
				leaveDate := listEnagedStatus[sub_index][1]
				// debug
				if leaveDate < "2019-09-23" {
					leaveDate = "2019-11-28"
				}
				avaliableSingle := calcAvaliables(startDate, endDate, arriveDate, leaveDate)
				avaliableUnmerged = append(avaliableUnmerged, avaliableSingle)
			}
			if len(avaliableUnmerged) == 0 {
				cntDays, _ := strconv.Atoi(endDate[8:10])
				for idxArray := 0; idxArray < cntDays; idxArray++ {
					avaliableMerged = append(avaliableMerged, 0)
				}
			} else {
				for idxAvaliablesUnmerged := 0; idxAvaliablesUnmerged < len(avaliableUnmerged[0]); idxAvaliablesUnmerged++ {
					statusMerged := -100
					for idx := 0; idx < len(avaliableUnmerged); idx++ {
						if statusMerged < avaliableUnmerged[idx][idxAvaliablesUnmerged] {
							statusMerged = avaliableUnmerged[idx][idxAvaliablesUnmerged]
						}
					}
					avaliableMerged = append(avaliableMerged, statusMerged)
				}
			}
			avaliables = append(avaliables, avaliableMerged)
		}
		jsonKutiInfo["avaliables"] = avaliables
		jsonKutiInfo["residents"] = listResidentsInfo
		jsonKutiInfo["unconfirmeds"] = unconfirmeds
		kutiesInfo = append(kutiesInfo, jsonKutiInfo)
	}
	// 当前及接下来四个月份
	currMonth, _ := strconv.Atoi(time.Now().Format("01"))
	currMonth -= 1
	months := [5]int{
		currMonth,
		(currMonth + 1) % 12,
		(currMonth + 2) % 12,
		(currMonth + 3) % 12,
		(currMonth + 4) % 12}
	retJson["months"] = months
	retJson["kutiesInfo"] = kutiesInfo
	// 待确认事件
	daySessions := getDateSessions()
	cnt1stMonthDays, _ := strconv.Atoi(daySessions[0][1][8:10])
	cnt2ndMonthDays, _ := strconv.Atoi(daySessions[1][1][8:10])
	cnt3rdMonthDays, _ := strconv.Atoi(daySessions[2][1][8:10])
	arrayMonthDaysCount := [3]int{cnt1stMonthDays, cnt2ndMonthDays, cnt3rdMonthDays}
	arrayToConfirmed := [3]([]bool){}
	arrayToConfirmed[0] = []bool{}
	arrayToConfirmed[1] = []bool{}
	arrayToConfirmed[2] = []bool{}
	for i := 0; i < 3; i++ {
		for j := 0; j < arrayMonthDaysCount[i]; j++ {
			hasTodo := false
			if j == 18 {
				hasTodo = true
			}
			arrayToConfirmed[i] = append(arrayToConfirmed[i], hasTodo)
		}
	}
	retJson["eventsToConfirm"] = arrayToConfirmed
	return retJson
}

func UpdateBrokenStatus(kutiNumber int, kutiType int, forSex int, brokenStatus int) bool {
	qs := o.QueryTable("tb_kuti").
		Filter("number", kutiNumber).
		Filter("type", kutiType).
		Filter("for_sex", forSex)
	var kuti Kuti
	err := qs.One(&kuti)
	if nil != err {
		println(err.Error())
		return false
	}
	kuti.Broken = brokenStatus
	cnt, error := o.Update(&kuti, "broken")
	if nil != error {
		println(error.Error())
	}
	println(fmt.Sprintf("effected row: %d", cnt))
	return true
}

func remainUnconfirmed(residentID int) bool {
	currDate := time.Now().Format("2006-01-02")
	items := []*Item{}
	query := o.QueryTable("tb_item").
		Filter("resident_id", residentID).
		Filter("enabled", 1).Filter("confirmed", 0).
		Filter("type__in", TYPE_APPOINT_TO_ARRIVE, TYPE_PLAN_TO_LEAVE).
		Filter("activate_date__lt", currDate)
	cnt, err := query.All(&items)
	if nil != err {
		println(err.Error())
		return false
	}
	return cnt != 0
}
