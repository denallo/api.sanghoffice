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
