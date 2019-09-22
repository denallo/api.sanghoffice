package models

import "fmt"

func GetKuties(forSex int) []map[string]interface{} {
	// 孤邸属性
	mapKuties := map[int]*Kuti{}
	var kuties []*Kuti
	query := o.QueryTable("tb_kuti").Filter("for_sex", forSex)
	num, err := query.All(&kuties)
	if nil != err {
		panic(err)
	}
	for i := 0; i < int(num); i++ {
		index := kuties[i].Id
		mapKuties[index] = kuties[i]
	}
	// 人员属性
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
	// 孤邸-人员映射状态
	var resiStatusList []*ResiStatus
	// mapResident2Status := map[int]*ResiStatus{}
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
		// mapResident2Status[resiStatus.ResidentId] = resiStatus
		mapKuti2Status[resiStatus.KutiId] = append(mapKuti2Status[resiStatus.KutiId], resiStatus)
	}
	// 打包json
	var retJson [](map[string]interface{})
	for i := 0; i < len(kuties); i++ {
		kutiInfo := kuties[i]
		item := map[string]interface{}{}
		item["KutiNumber"] = kutiInfo.Number
		var listResidentsInfo [](map[string]interface{})
		for j := 0; j < len(mapKuti2Status[kutiInfo.Id]); j++ {
			residentInfo := map[string]interface{}{}
			resiStatus := mapKuti2Status[kutiInfo.Id][j]
			resident, existed := mapResidents[resiStatus.ResidentId]
			if !existed {
				fmt.Println(resiStatus.ResidentId)
				continue
			}
			residentInfo["Id"] = resident.Id
			if R_TYPE_BHIKHU == resident.Type ||
				R_TYPE_SAMANERA == resident.Type ||
				R_TYPE_SAYALAY == resident.Type ||
				R_TYPE_OTHER_MONK == resident.Type {
				residentInfo["Name"] = resident.Dhamame
				residentInfo["IsMonk"] = 1
			} else {
				residentInfo["Name"] = resident.Name
				residentInfo["IsMonk"] = 0
			}
			listResidentsInfo = append(listResidentsInfo, residentInfo)
		}
		item["Residents"] = listResidentsInfo
		retJson = append(retJson, item)
	}
	return retJson
}
