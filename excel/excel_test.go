package excel

import "testing"

func TestNewCommonExcel(t *testing.T) {

	e := NewCommonExcel("", "test")

	e.WriteData(1, []interface{}{1, 2, 3, 4})
	e.WriteData(2, []interface{}{5, 6, 7, 8})
	e.Save()
}
