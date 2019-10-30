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

func AddResiStatus(residentID int, sex int, kutiNumber int, kutiType int, arriveDate string, leaveDate string) bool {
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

func GetAvailablesInfo(kutiNumber int, kutiType int, forSex int) ([]([]int), bool) {
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
		enagedInfo := [2]string{resiStatus.ArriveDate, resiStatus.PlanToLeaveDate}
		listEnagedStatus = append(listEnagedStatus, enagedInfo)
	}

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
	return avaliables, true
}

func CreateItem(residentID int, arriveDate string, leaveDate string) bool {
	items := []*Item{}
	query := o.QueryTable("tb_item").Filter("resident_id", residentID).Filter("confirmed", 0)
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
