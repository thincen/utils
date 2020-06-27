// 方糖 http://sc.ftqq.com/
// 推送消息到微信的工具

package ft

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// FT client
type FT struct {
	sckey string
	url   string
}

// NewClient 创建对象
func NewClient(key string) *FT {
	return &FT{
		sckey: key,
		url:   "https://sc.ftqq.com/" + key + ".send",
	}
}

// Send msg, text:subject; desp:content
// 返回值.成功:{"errno":0,"errmsg":"success","dataset":"done"};失败:{"errno":1024,"errmsg":"bad pushtoken"}
func (ft *FT) Send(text, desp string) error {
	var v = make(url.Values, 2)
	v.Add("text", text)
	v.Add("desp", desp)
	res, err := http.PostForm(ft.url, v)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	r := &struct {
		Errno   int    `json:"errno"`
		Errmsg  string `json:"errmsg"`
		Dataset string `json:"dataset"`
	}{}
	err = json.Unmarshal(buf, r)
	if err != nil {
		return err
	}
	if r.Errno != 0 {
		return errors.New(r.Errmsg)
	}
	return nil
}
