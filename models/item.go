package models

import (
	"fmt"
	"strconv"
)

func GetBrief(year int, month int) ([]int, bool) {
	pattern := fmt.Sprintf("%d-%02d-%%", year, month)
	items := []*Item{}
	cnt, err := o.Raw("SELECT * FROM tb_item WHERE activate_date like ?", pattern).QueryRows(&items)
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
