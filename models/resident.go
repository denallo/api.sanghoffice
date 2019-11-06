package models

import (
	"api.sanghoffice/tools"
	"github.com/astaxie/beego/orm"
)

func IsExistedResident(name string, isDhamame bool, sex int) (existed bool, residentID int) {
	var resident Resident
	key := ""
	if isDhamame {
		resident.Dhamame = name
		key = "dhamame"
	} else {
		resident.Name = name
		key = "name"
	}
	resident.Sex = sex
	if orm.ErrNoRows == o.Read(&resident, key) {
		existed = false
	} else {
		existed = true
		residentID = resident.Id
	}
	return existed, residentID
}

func GetResidents(sex int) ([]ResidentInTemple, bool) {
	var residents []ResidentInTemple
	_, err := o.Raw("SELECT * from v_resident_in_temple WHERE sex = ?", sex).QueryRows(&residents)
	if err != nil {
		println(err.Error())
	}
	return residents, true
}

func AddResident(data map[string]interface{}) (residentID int) {
	resident := Resident{}
	for key, value := range data {
		if key == "name" {
			resident.Name = value.(string)
		} else if key == "dhamame" {
			resident.Dhamame = value.(string)
		} else if key == "sex" {
			resident.Sex, _ = tools.JsonNumberToInt(value)
		} else if key == "identifier" {
			resident.Identifier = value.(string)
		} else if key == "age" {
			resident.Age, _ = tools.JsonNumberToInt(value)
		} else if key == "type" {
			resident.Type, _ = tools.JsonNumberToInt(value)
		} else if key == "folk" {
			resident.Folk = value.(string)
		} else if key == "native_place" {
			resident.NativePlace = value.(string)
		} else if key == "ability" {
			resident.Ability = value.(string)
		} else if key == "phone" {
			resident.Phone = value.(string)
		} else if key == "emergency_contact" {
			resident.EmergencyContact = value.(string)
		} else if key == "emergency_contact_phone" {
			resident.EmergencyContactPhone = value.(string)
		}
	}
	id, err := o.Insert(&resident)
	if err != nil {
		println(err)
		return -1
	}
	residentID = int(id)
	return residentID
}
