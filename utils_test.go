package utils

import (
	"runtime"
	"testing"
)

func Test_GetFucNameLine(t *testing.T) {
	t.Log(GetFuncNameLine(1))
	// utils_test.go-Test_GetFucNameLine:9
}

func Test_GenerAbsPath(t *testing.T) {
	abspath := "/home/thincen/.vimrc"
	if runtime.GOOS == "windows" {
		abspath = "D:\\App\\test.txt"
	}
	relpath := "utils.go"
	t.Log(GenerAbsPath(abspath))
	t.Log(GenerAbsPath(relpath))
}

func Test_IsExist(t *testing.T) {
	if IsExist("utils.go") {
		t.Log("pass")
	} else {
		t.Log("fail")
	}
}
