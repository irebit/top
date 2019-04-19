package utils

import (
	"encoding/json"
	"strconv"
	"time"
)

var (
	TIME_LAYOUT = "2006-01-02 15:04:05"
	TIME_ZONE   = "Asia/Shanghai"
)

func GetCurrentDateTime() string {
	loc, _ := time.LoadLocation(TIME_ZONE)
	ti := time.Now().In(loc)
	return ti.Format(TIME_LAYOUT)
}

func GetString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	case byte:
		return string(v)
	default:
		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
}
