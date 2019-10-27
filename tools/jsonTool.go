package tools

import "encoding/json"

func JsonNumberToInt(jsField interface{}) (int, bool) {
	interger64, err := jsField.(json.Number).Int64()
	if err != nil {
		println(err.Error())
		return -1, false
	}
	interger32 := int(interger64)
	return interger32, true
}
