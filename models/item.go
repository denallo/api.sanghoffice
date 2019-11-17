package models

import (
	"fmt"
	"strconv"
)

func GetBrief(year int, month int) ([]int, bool) {
	pattern := fmt.Sprintf("%d-%02d-%%", year, month)
	items := []*Item{}
	cnt, err := o.Raw(
		"SELECT * FROM tb_item "+
			"WHERE activate_date like ? "+
			"AND confirmed = 0",
		pattern).QueryRows(&items)
	if err != nil {
		println(err.Error())
		return nil, false
	}
	result := []int{}
	for i := 0; i < 31; i++ {
		result = append(result, 0)
	}
	for i := 0; i < int(cnt); i++ {
		activateDate := items[i].ActivateDate
		dayStr := activateDate[8:len(activateDate)]
		day, _ := strconv.Atoi(dayStr)
		result[day-1] += 1
	}
	return result, true
}

func UpdateResidentState(residentID int, stateType int) bool {
	sql := o.Raw("call proc_confirm_item(?, ?)", residentID, stateType)
	success := -1
	err := sql.QueryRow(&success)
	if nil != err {
		println(err.Error())
		return false
	} else if success != 0 {
		return false
	}
	return true
}
