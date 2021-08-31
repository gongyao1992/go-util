package go_util

import "encoding/json"

// JsonToStruct Json数组变为结构体
func JsonToStruct(i interface{}, toStruct interface{}) error {
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

// ToArr 字符串变为 结构体
func ToArr(jstr string, toStruct interface{}) error {
	return json.Unmarshal([]byte(jstr), &toStruct)
}

// ExcelLie 通过传入一个数字 获取excel对应的 列
func ExcelLie(i int) string {
	lie1Number := getLie1Number(i)
	lie2Number := getLie2Number(i)
	return excelLie(rune(lie1Number)) + excelLie(rune(lie2Number))
}

func excelLie(i rune) string {
	if i <= 0 {
		return ""
	}
	b := []rune{64}
	b[0] += i

	return string(b)
}

func getLie1Number(x int) (y int) {
	if x % 26 == 0 { // 整除
		y = (x / 26) - 1
	} else {
		y = x / 26
	}
	return
}

func getLie2Number(x int) (y int) {
	if x % 26 == 0 { // 整除
		y = 26
	} else {
		y = x % 26
	}
	return
}