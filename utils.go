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

// GenerAbsPath 如果path是相对地址则转换为binary文件同级目录地址
// 如果是绝对路径则返回path
// 错误返回path和error
func GenerAbsPath(path string) (string, error) {
	var res string = ""
	info, e := os.Lstat(os.Args[0])
	if e != nil {
		return path, e
	}
	// 链接
	if info.Mode()&os.ModeSymlink != 0 {
		binpath, e := filepath.EvalSymlinks(os.Args[0])
		if e != nil {
			return path, e
		}
		res = filepath.Dir(binpath) + "/" + path
	} else {
		// 文件
		if filepath.IsAbs(path) {
			return path, nil
		}
		abspath, e := filepath.Abs(os.Args[0])
		if e != nil {
			return path, e
		}
		res = filepath.Dir(abspath) + "/" + path
	}
	// Unix/Windows support
	if runtime.GOOS == "windows" {
		res = strings.ReplaceAll(res, "/", "\\")
	}
	return res, nil
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
