package models

import (
	"fmt"

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

func GetResidentInfo(residentId int) (Resident, bool) {
	info := Resident{Id: residentId}
	err := o.Read(&info)
	if err != nil {
		println(err.Error())
		return info, false
	}
	return info, true
}

// const TYPE_RESIDENT = 0
// const TYPE_PLAN_TO_LEAVE = 1
// const TYPE_PLAN_TO_LEAVE_CONFIRMED = 2
// const TYPE_APPOINT_TO_ARRIVE = 3
// const TYPE_APPOINT_TO_ARRIVE_CONFIRMED = 4

const PRE_HINT_DAYS = 3

func GetResidents(sex int, state int) ([]ResidentInTemple, bool) {
	var residents []ResidentInTemple
	// sqlCond := ""
	// sqlCurrDate := ""
	// sqlPreHintDate := ""
	// switch state {
	// case TYPE_RESIDENT:
	// 	sqlCond = "`arrive` <= DATE_FORMAT(NOW(),'%Y-%m-%d') AND `leave` > DATE_FORMAT(NOW(),'%Y-%m-%d')"
	// 	break
	// case TYPE_APPOINT_TO_ARRIVE:
	// 	sqlCurrDate = "DATE_FORMAT(NOW(), '%Y-%m-%d')"
	// 	sqlPreHintDate = fmt.Sprintf(
	// 		"DATE_ADD(STR_TO_DATE(`arrive`, '%%Y-%%m-%%d'), interval -%d day)",
	// 		PRE_HINT_DAYS)
	// 	sqlCond = fmt.Sprintf(
	// 		"(%s >= %s AND %s <= `arrive`)",
	// 		sqlCurrDate, sqlPreHintDate,
	// 		sqlCurrDate)
	// 	break
	// case TYPE_PLAN_TO_LEAVE:
	// 	sqlCurrDate = "DATE_FORMAT(NOW(), '%Y-%m-%d')"
	// 	sqlPreHintDate = fmt.Sprintf(
	// 		"DATE_ADD(STR_TO_DATE(`leave`, '%%Y-%%m-%%d'), interval -%d day)",
	// 		PRE_HINT_DAYS)
	// 	sqlCond = fmt.Sprintf(
	// 		"(%s >= %s AND %s <= `leave`)",
	// 		sqlCurrDate, sqlPreHintDate,
	// 		sqlCurrDate)
	// 	break
	// default:
	// 	break
	// }
	// sql := fmt.Sprintf("SELECT * FROM v_resident_in_temple WHERE sex = %d AND %s", sex, sqlCond)
	// sqlCurrDate = "DATE_FORMAT(NOW(), '%Y-%m-%d')"
	// subSql := fmt.Sprintf("SELECT resident_id FROM tb_item WHERE type=%d AND enabled=1 AND ()", state)
	sql := fmt.Sprintf(
		"SELECT * FROM v_residents WHERE resident_id IN ("+
			"SELECT resident_id FROM tb_item WHERE type=%d "+
			"AND enabled = 1 AND confirmed != 1 "+
			"AND (activate_date = '' OR activate_date <= DATE_FORMAT(NOW(), '%%Y-%%m-%%d')))",
		state)
	if state == TYPE_APPOINTED {
		sql = fmt.Sprintf(
			"SELECT * FROM v_residents WHERE resident_id IN (" +
				"SELECT resident_id FROM tb_item WHERE type = 0 " +
				"AND enabled = 1 AND confirmed != 1)")
	}
	println(sql)
	_, err := o.Raw(sql).QueryRows(&residents)
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

func UpdateResident(data map[string]interface{}) (Resident, bool) {
	id, success := tools.JsonNumberToInt(data["id"])
	if !success {
		return Resident{}, false
	}
	resident := Resident{Id: id}
	err := o.Read(&resident, "id")
	if err != nil {
		println(err.Error())
		return Resident{}, false
	}
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
	o.Update(&resident)
	resident = Resident{Id: id}
	err = o.Read(&resident, "id")
	if err != nil {
		println(err.Error())
		return Resident{}, false
	}
	return resident, true
}
