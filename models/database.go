package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

type Kuti struct {
	Id     int `orm:"column(id)"`
	Number int `orm:"column(number)"`
	Type   int `orm:"column(type)"`
	ForSex int `orm:"column(for_sex)"`
	Broken int `orm:"column(broken)"` // 0-正常 1-损坏
}

func (tb *Kuti) TableName() string {
	return "tb_kuti"
}

type Resident struct {
	Id                    int    `orm:"column(id)"`
	Name                  string `orm:"column(name)"`
	Dhamame               string `orm:"column(dhamame)"`
	Sex                   int    `orm:"column(sex)"`
	Identifier            string `orm:"column(identifier)"`
	Age                   int    `orm:"column(age)"`
	Type                  int    `orm:"column(type)"`
	Folk                  string `orm:"column(folk)"`
	NativePlace           string `orm:"column(native_place)"`
	Ability               string `orm:"column(ability)"`
	Phone                 string `orm:"column(phone)"`
	EmergencyContact      string `orm:"column(emergency_contact)"`
	EmergencyContactPhone string `orm:"column(emergency_contact_phone)"`
}

const (
	R_TYPE_MONK_UNCERTAIN = 100 // 出家众,未确认具体类型
	R_TYPE_BHIKHU         = 0
	R_TYPE_SAMANERA       = 1
	R_TYPE_SAYALAY        = 2
	R_TYPE_OTHER_MONK     = 3
	R_TYPE_SERVER         = 4
	R_TYPE_DHAMWORKER     = 5
	R_TYPE_OTHER_HOMER    = 6
)

func (tb *Resident) TableName() string {
	return "tb_resident"
}

type ResiStatus struct {
	ResidentId      int    `orm:"column(resident_id);pk"`
	KutiId          int    `orm:"column(kuti_id)"`
	ArriveDate      string `orm:"column(arrive_date)"`
	PlanToStayDays  int    `orm:"column(plan_to_stay_days)"`
	PlanToLeaveDate string `orm:"column(plan_to_leave_date)"`
	TurnedPhoneCard int    `orm:"column(turned_phone_card)"`
}

func (tb *ResiStatus) TableName() string {
	return "tb_resi_status"
}

type ResiHistory struct {
	Id         int    `orm:"column(id)"`
	ResidentId int    `orm:"column(resident_id)"`
	KutiId     int    `orm:"column(kuti_id)"`
	ArriveDate string `orm:"column(arrive_date)"`
	LeaveDate  string `orm:"column(leave_date)"`
	Comment    string `orm:"column(comment)"`
}

func (tb *ResiHistory) TableName() string {
	return "tb_resi_history"
}

type Instance struct {
	Identifier string `orm:"pk;column(identifier)"`
	ClassName  string `orm:"column(class_name)"`
	Property   string `orm:"column(property)"`
	Value      string `orm:"column(value)"`
}

func (tb *Instance) TableName() string {
	return "tb_instances"
}

type Relations struct {
	Id          int64
	IdentifierA string `orm:"column(identifier_a)"`
	IdentifierB string `orm:"column(identifier_b)"`
	ClassNameA  string `orm:"column(class_name_a)"`
	ClassNameB  string `orm:"column(class_name_b)"`
	RelId       string `orm:"column(rel_id)"`
	Describe    string `orm:"column(describe)"`
}

func (tb *Relations) TableName() string {
	return "tb_relations"
}

func DoSync() {
	// instance-kuti
	var kuties []*Instance
	kutiID := 1
	residentID := 1
	query := o.QueryTable("tb_instances").Filter("class_name", "kuti")
	_, _ = query.All(&kuties)
	mapKutiInstance := map[string]Kuti{}
	mapKutiId := map[string]int{}
	for i := 0; i < len(kuties); i++ {
		id := kuties[i].Identifier
		if _, existed := mapKutiInstance[id]; !existed {
			mapKutiInstance[id] = Kuti{}
			mapKutiId[id] = kutiID
			kutiID += 1
		}
		kuti := mapKutiInstance[id]
		property := kuties[i].Property
		value := kuties[i].Value
		if property == "number" {
			kuti.Number, _ = strconv.Atoi(value)
		} else if property == "for_sex" {
			kuti.ForSex, _ = strconv.Atoi(value)
		} else if property == "type" {
			kuti.Type, _ = strconv.Atoi(value)
		}
		mapKutiInstance[id] = kuti
	}
	var kutiRecords []Kuti
	for key := range mapKutiInstance {
		kutiRecords = append(kutiRecords, mapKutiInstance[key])
	}
	num, err := o.InsertMulti(len(kutiRecords), kutiRecords)
	fmt.Println(num, err)
	// instance-resident
	var residents []*Instance
	query = o.QueryTable("tb_instances").Filter("class_name", "resident").Limit(-1)
	_, _ = query.All(&residents)
	// fmt.Println(num, err)
	mapResidentInstance := map[string]Resident{}
	mapResiStatus := map[string]ResiStatus{}
	mapResidentId := map[string]int{}
	for i := 0; i < len(residents); i++ {
		id := residents[i].Identifier
		if _, existed := mapResidentInstance[id]; !existed {
			mapResidentInstance[id] = Resident{}
			mapResidentId[id] = residentID
			residentID += 1
		}
		if _, existed := mapResiStatus[id]; !existed {
			mapResiStatus[id] = ResiStatus{}
		}
		resident := mapResidentInstance[id]
		resiStatus := mapResiStatus[id]
		property := residents[i].Property
		value := residents[i].Value
		if property == "ability" {
			resident.Ability = value
		} else if property == "age" {
			resident.Age, _ = strconv.Atoi(value)
		} else if property == "dhamame" {
			resident.Dhamame = value
		} else if property == "emergency_contact" {
			resident.EmergencyContact = value
		} else if property == "emergency_contact_phone" {
			resident.EmergencyContactPhone = value
		} else if property == "Folk" {
			resident.Folk = value
		} else if property == "identifier" {
			resident.Identifier = value
		} else if property == "name" {
			resident.Name = value
		} else if property == "native_place" {
			resident.NativePlace = value
		} else if property == "phone" {
			resident.Phone = value
		} else if property == "sex" {
			resident.Sex, _ = strconv.Atoi(value)
		} else if property == "type" {
			resident.Type, _ = strconv.Atoi(value)
		} else if property == "prepare_leave_date" {
			resiStatus.PlanToLeaveDate = value
		} else if property == "residence" {
			resiStatus.PlanToStayDays, _ = strconv.Atoi(value)
		} else if property == "turned_in" {
			turnedPhoneCard, err := strconv.Atoi(value)
			if err != nil {
				resiStatus.TurnedPhoneCard = turnedPhoneCard
			}
		}
		mapResidentInstance[id] = resident
		mapResiStatus[id] = resiStatus
	}
	var residentRecords []Resident
	for key := range mapResidentInstance {
		residentRecords = append(residentRecords, mapResidentInstance[key])
	}
	num, err = o.InsertMulti(len(residentRecords), residentRecords)
	fmt.Println(num, err)
	// relations
	var relations []*Relations
	// mapResiStatus := map[string]ResiStatus{}
	query = o.QueryTable("tb_relations").Filter("class_name_a", "kuti").Limit(-1)
	num, err = query.All(&relations)
	for i := 0; i < len(relations); i++ {
		idResident := relations[i].IdentifierB
		idKuti := relations[i].IdentifierA
		resiStatus := mapResiStatus[idResident]
		resiStatus.KutiId = mapKutiId[idKuti]
		resiStatus.ResidentId = mapResidentId[idResident]
		if resiStatus.KutiId == 0 || resiStatus.ResidentId == 0 {
			// fmt.Println("ouch")
			delete(mapResiStatus, idResident)
			continue
		}
		mapResiStatus[idResident] = resiStatus
	}
	var lstResiStatus []ResiStatus
	for idResident := range mapResiStatus {
		resiStatus := mapResiStatus[idResident]
		if resiStatus.KutiId == 0 || resiStatus.ResidentId == 0 {
			continue
		}
		lstResiStatus = append(lstResiStatus, mapResiStatus[idResident])
	}
	num, err = o.InsertMulti(len(lstResiStatus), lstResiStatus)
	fmt.Println(num, err)
}

func init() {
	orm.RegisterDataBase("default", "mysql", "sanghoffice:fzwl2019@tcp(localhost:3306)/sanghoffice?charset=utf8")
	orm.RegisterModel(new(Kuti))
	orm.RegisterModel(new(Resident))
	orm.RegisterModel(new(ResiStatus))
	orm.RegisterModel(new(ResiHistory))
	// orm.RegisterModel(new(Instance))
	// orm.RegisterModel(new(Relations))
	orm.RunSyncdb("default", false, false)
	o = orm.NewOrm()
	orm.Debug = true
	// DoSync()
}
