package ft

import "testing"

func Test_NewClient(t *testing.T) {
	ft := NewClient("testsckey")
	t.Log(ft)
}

func Test_Send(t *testing.T) {
	ft := NewClient("SCUxxxxx")
	text := `FTsdk`
	desp := `# ft_test.go

> github.com/thincen/utils

- ft
- crypto aes base64
- other

~~email~~: **no_1seed@163.com**
`
	t.Log(ft.Send(text, desp))
}
