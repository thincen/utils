package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//IsExist 检查文件是否存在
func IsExist(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//GenerAbsPath 如果path是相对地址则转换为binary文件同级目录地址
//如果是绝对路径则返回path
func GenerAbsPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	binpath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// Unix/Windows support
	if runtime.GOOS == "windows" {
		path = binpath + "\\" + strings.ReplaceAll(path, "/", "\\")
	} else {
		path = binpath + "/" + path
	}
	return path
}

// GetFuncNameLine 运行到哪一行
func GetFuncNameLine(skip int) string {
	fn, file, line, ok := runtime.Caller(skip)
	if ok {
		funcName := strings.Split(runtime.FuncForPC(fn).Name(), ".")
		f := strings.Split(file, "/")
		return fmt.Sprintf("%s-%s:%d ", f[len(f)-1], funcName[len(funcName)-1], line)
	}
	return ""
}
