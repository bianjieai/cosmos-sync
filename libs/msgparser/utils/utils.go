package utils

import (
	"encoding/json"
)

func MarshalJsonIgnoreErr(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

func UnMarshalJsonIgnoreErr(data string, v interface{}) {
	json.Unmarshal([]byte(data), &v)
}
