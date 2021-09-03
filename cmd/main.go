package main

import "github.com/gongyao1992/go-util/excel"

func main()  {
	e := excel.NewCommonExcel("", "test")

	e.WriteData(1, []interface{}{1, 2, 3, 4})
	e.WriteData(2, []interface{}{5, 6, 7, 8})
	e.Save()
}
