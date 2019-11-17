package models

import (
	"fmt"
	"strconv"
)

type ReqNewResiStatus struct {
	ArriveDate            string `json:"arriveDate"`
	LeaveDate             string `json:"leaveDate"`
	Ability               string `json:"ability"`
	Age                   int    `json:"age"`
	Dhamame               string `json:"dhamame"`
	EmergencyContact      string `json:"emergency_contact"`
	EmergencyContactPhone string `json:"emergency_contact_phone"`
	Folk                  string `json:"folk"`
	Identifier            string `json:"identifier"`
	Name                  string `json:"name"`
	NativePlace           string `json:"native_place"`
	Phone                 string `json:"phone"`
	Sex                   int    `json:"sex"`
}

func AddResiStatus(
	residentID int,
	sex int,
	kutiNumber int,
	kutiType int,
	arriveDate string,
	leaveDate string) bool {

	kuti := Kuti{Number: kutiNumber, ForSex: sex, Type: kutiType}
	err := o.Read(&kuti, "number", "for_sex", "type")
	if err != nil {
		println(err.Error())
		return false
	}
	resiStatus := ResiStatus{
		ResidentId:      residentID,
		KutiId:          kuti.Id,
		ArriveDate:      arriveDate,
		PlanToLeaveDate: leaveDate,
	}
	_, error := o.Insert(&resiStatus)
	if error != nil {
		println(error.Error())
		return false
	}
	return true
}

func GetAvailablesInfo(
	kutiNumber int,
	kutiType int,
	forSex int) ([]([]int), bool) {

	kuti := Kuti{Number: kutiNumber, ForSex: forSex, Type: kutiType}
	error := o.Read(&kuti, "number", "for_sex", "type")
	if nil != error {
		println(error.Error())
		return nil, false
	}
	kutiID := kuti.Id

	resiStatusList := []*ResiStatus{}
	query := o.QueryTable("tb_resi_status").Filter("kuti_id", kutiID)
	cnt, err := query.All(&resiStatusList)
	if nil != err {
		println(err.Error())
		return nil, false
	}

	var listEnagedStatus []([2]string) // [(arriveDate, leaveDate), ...]
	for i := 0; i < int(cnt); i++ {
		resiStatus := resiStatusList[i]
		enagedInfo := [2]string{
			resiStatus.ArriveDate,
			resiStatus.PlanToLeaveDate}
		listEnagedStatus = append(listEnagedStatus, enagedInfo)
	}

	dateSessions := getDateSessions()
	var avaliables []([]int)
	for i := 0; i < len(dateSessions); i++ {
		// 用于合并一栋孤邸多个住众在同一天的入住状态
		var avaliableUnmerged []([]int)
		var avaliableMerged []int
		startDate := dateSessions[i][0]
		endDate := dateSessions[i][1]
		for j := 0; j < len(listEnagedStatus); j++ {
			arriveDate := listEnagedStatus[j][0]
			leaveDate := listEnagedStatus[j][1]
			avaliableSingle := calcAvaliables(
				startDate, endDate,
				arriveDate, leaveDate)
			avaliableUnmerged = append(avaliableUnmerged, avaliableSingle)
		}
		if len(avaliableUnmerged) == 0 {
			cntDays, _ := strconv.Atoi(endDate[8:10])
			for idxArray := 0; idxArray < cntDays; idxArray++ {
				avaliableMerged = append(avaliableMerged, 0)
			}
		} else {
			for j := 0; j < len(avaliableUnmerged[0]); j++ {
				statusMerged := -100
				for k := 0; k < len(avaliableUnmerged); k++ {
					if statusMerged < avaliableUnmerged[k][j] {
						statusMerged = avaliableUnmerged[k][j]
					}
				}
				avaliableMerged = append(avaliableMerged, statusMerged)
			}
		}
		avaliables = append(avaliables, avaliableMerged)
	}
	return avaliables, true
}

func CreateItem(
	residentID int,
	arriveDate string,
	leaveDate string) bool {

	items := []*Item{}
	query := o.QueryTable("tb_item").
		Filter("resident_id", residentID).
		Filter("confirmed", 0)
	cnt, err := query.All(&items)
	if err != nil {
		println(err.Error())
		return false
	}
	for i := 0; i < int(cnt); i++ {
		item := items[i]
		var activateDate string
		switch item.Type {
		case 0: // 人员于当天离开
			// t, _ := time.Parse(DATE_LAYOUT, leaveDate)
			// before, _ := time.ParseDuration("-72h")
			// activateDate = t.Add(before).Format(DATE_LAYOUT)
			activateDate = leaveDate
			break
		case 1: // 已预约人员于当天到达
			// t, _ := time.Parse(DATE_LAYOUT, arriveDate)
			// before, _ := time.ParseDuration("-72h")
			// activateDate = t.Add(before).Format(DATE_LAYOUT)
			activateDate = arriveDate
			break
		default:
			println(fmt.Sprintf("unknow item type: %d", item.Type))
			return false
		}
		item.ActivateDate = activateDate
		_, err := o.Update(item, "activate_date")
		if err != nil {
			println(err.Error())
			return false
		}
	}
	return true
}

func CheckIn(
	name string,
	dhamame string,
	identifer string,
	sex int,
	age int,
	residentType int,
	folk string,
	nativePlace string,
	ability string,
	phone string,
	emergencyContact string,
	emergencyContactPhone string,
	kutiNumber int,
	kutiType int,
	isMonk int,
	arriveDate string,
	leaveDate string) (int, bool) {

	sql := o.Raw(
		"call proc_check_in("+
			"?, ?, ?, ?, ?,"+
			"?, ?, ?, ?, ?,"+
			"?, ?, ?, ?, ?,"+
			"?, ?)",
		name, dhamame, identifer, sex,
		age, residentType, folk,
		nativePlace, ability, phone,
		emergencyContact, emergencyContactPhone,
		kutiNumber, kutiType, isMonk,
		arriveDate, leaveDate)
	residentID := -1
	err := sql.QueryRow(&residentID)
	if nil != err {
		println(err.Error())
		return -1, false
	}
	return residentID, true
}
