# package ft

为方便Golang中调用方糖服务创建。

## 方糖介绍

方糖推送服务 ([Server酱](http://sc.ftqq.com/)) 是一款「程序员」和「服务器」之间的通信软件。

1. 登入：用GitHub账号登入网站，就能获得一个SCKEY（在「[发送消息](http://sc.ftqq.com/?c=code)」页面）
2. 绑定：点击「微信推送」，扫码关注同时即可完成绑定
3. 发消息：往 `http://sc.ftqq.com/SCKEY.send` 发GET请求，就可以在微信里收到消息啦

## Usage

```shell
# install
go get -u github.com/thincen/utils/ft
```

### Example

```golang
// demo.go

package main

import (
    "fmt"

    "github.com/thincen/utils/ft"
)

func main(){
    const sckey string = "SUCxxxx"
    ft := ft.NewClient(sckey)
	text := `FTsdk`
	desp := `# ft_test.go

> github.com/thincen/utils

- ft
- crypto aes+base64
- other

~~email~~: **no_1seed@163.com**
`
	if err := ft.Send(text, desp);err!=nil{
        fmt.Println(err)
    }
}
```
