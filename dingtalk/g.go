package dingtalk

import (
	"fmt"
	"github.com/CatchZeng/dingtalk"
	"strconv"
)


func DingtalkSend2()  {
	accessToken := "a071c2eaa8538a4a70a953556e6bc9acc9aa1248b6bcd7534c3e5672fa848fc4"
	secret := "SEC52ba5288b2945901f4986f474210f86e091547585d89a9467cf7675f2c204965"

	client := dingtalk.NewClient(accessToken, secret)


	//s := syncTable{
	//	TableName:      "1",
	//	DumpBeginTime:  "2",
	//	DumpEndTime:    "3",
	//	DumpErrorMsg:   "4",
	//	DumpTime:       2,
	//	FileName:       "",
	//	ImportEndTime:  "",
	//	ImportErrorMsg: "",
	//	ImportTime:     0,
	//}


	s := make([]string, 0)
	i := 0
	for true {
		s = append(s, strconv.Itoa(i))
		i ++
		if i >= 1000 {
			break
		}
	}
	//msg := dingtalk.NewTextMessage().SetContent(strings.Join(s, ",")).SetAt([]string{"15010062048"}, false)
	//msg := dingtalk.NewMarkdownMessage().SetMarkdown("111",strings.Join(s, ",") ).SetAt([]string{"15010062048"}, false)
	//picUrl := "https://files.56hello.cn/9275-1636705742-9401.jpeg?e=1636709775&token=jU4wipkkaxqfrcsgyZ2lcrNuAiXNIu1WPZTwxAkw:4Thr9o7LL8nospHnkcuLicantc0="
	picUrl := ""
	msg := dingtalk.NewLinkMessage().SetLink("title", "text", picUrl, "https://github.com/")
	rep, err := client.Send(msg)
	//rep.ErrMsg
	fmt.Println(rep.ErrMsg, err)
}

type syncTable struct {
	TableName		string		`json:"table_name"`		// 表名
	DumpBeginTime 	string		`json:"dump_begin_time"` // 导出开始时间
	DumpEndTime		string		`json:"dump_end_time"` // 导出结束时间
	DumpErrorMsg	string 		`json:"dump_error_msg"` // 导出的错误信息
	DumpTime		float64 	`json:"dump_time"` // 导出时间
	FileName		string 		`json:"file_name"` // 文件名
	ImportEndTime	string 		`json:"import_end_time"` // 导入截止时间
	ImportErrorMsg	string 		`json:"import_error_msg"` // 导入错误信息
	ImportTime		float64 	`json:"import_time"` // 导入时间
}

func (s syncTable)GetFormat() string {
	//s.TableName
	return fmt.Sprintf("## 表名%s		\n导出执行开始时间：%s	\n  导出结束时间：%s	\n  导出的错误信息：%s	\n  导出的执行时间：%f	\n",
		s.TableName, s.DumpBeginTime, s.DumpEndTime, s.DumpErrorMsg, s.DumpTime)
}