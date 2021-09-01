package helper

import "encoding/json"

// InterfaceToStruct 接口数据变结构体
func InterfaceToStruct(i interface{}, toStruct interface{}) error {
	resByre, err := json.Marshal(i)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resByre, &toStruct)
	return err
}

// ToJson 一切结构体变为json 字符串
func ToJson(i interface{}) string {
	resByre, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(resByre)
}

// StrToStruct 字符串变为 结构体
func StrToStruct(jstr string, toStruct interface{}) error {
	return json.Unmarshal([]byte(jstr), &toStruct)
}