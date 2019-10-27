package models

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
