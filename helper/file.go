package helper

import (
	"os"
	"path/filepath"
	"strings"
)

// PathHasFile 路径下是否有文件
func PathHasFile(path, file string) bool {
	filepathNames, err := filepath.Glob(filepath.Join(path, file))
	if err != nil {
		return false
	}
	if len(filepathNames) == 0 {
		return false
	}
	return true
}

// GetConfigFile 获取配置文件的全称
func GetConfigFile(configFile string) string {
	// 获取工作空间
	wd, _ := os.Getwd()
	t := 1
	// 找到配置文件的路径
	for !PathHasFile(wd, configFile) {
		// 如果一直没有 那么最多校验6层
		if t >= 6 {
			panic("配置文件不存在")
		}

		wd += "/.." // 一直向上遍历
		t += 1
	}
	return wd + "/" + configFile
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetFileNameByPath 通过路径获取文件名
func GetFileNameByPath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr) - 1]
}