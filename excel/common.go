package excel

import (
	"github.com/gongyao1992/go-util/helper"
	excelize "github.com/xuri/excelize/v2"
	"os"
	"strconv"
)

// getExcelLie 通过传入一个数字 获取excel对应的 列
func getExcelLie(i int) string {
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


const ex = ".xlsx"

type CommonExcel struct {
	dir string
	fileName string
	file *excelize.File
}

func NewCommonExcel(dir, fileName string) *CommonExcel {
	f := excelize.NewFile()
	e := &CommonExcel{file: f, dir: dir}

	exists, _ := helper.PathExists(e.GetFile())
	if !exists {
		os.Mkdir(e.GetFile(), 0777)
	}

	if len(fileName) > 0 {
		e.fileName = fileName + ex
	}
	return e
}

func (e *CommonExcel)SetFileName(fileName string)  {
	e.fileName = fileName + ex
}

func (e *CommonExcel)WriteData(hangIndex int, values []interface{}) error {
	if len(values) == 0 {
		return nil
	}
	lie := 1
	for _, values := range values {
		err := e.file.SetCellValue("Sheet1", getZuobiao(hangIndex, lie), values)
		if err != nil {
			return err
		}
		lie += 1
	}

	return nil
}

func (e *CommonExcel)Save() error {
	return e.file.SaveAs(e.GetFile())
}

func (e *CommonExcel)GetFilePath() string {
	panic("del")
}
func (e *CommonExcel)GetFile() string {
	//d := "." + e.dir
	d := e.dir
	d += e.fileName
	return d
}

// GetFileName 获取文件名称
func (e *CommonExcel)GetFileName() string {
	return e.fileName
}

// 获取坐标
func getZuobiao(hangInt, lieInt int) string {
	lieStr := getExcelLie(lieInt)
	axls := lieStr + strconv.Itoa(hangInt)
	return axls
}