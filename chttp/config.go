//author: https://github.com/zhaojunlike
//date: 2019/12/12
package chttp

import (
	"time"
)

//请求配置
type Config struct {
	Url             string
	Headers         map[string]string
	Timeout         time.Duration
	Data            interface{}
	Method          string
	Retry           int
	ExecCount       int
	PostForm        bool
	AutoParseCookie bool
	ParseJson       bool
	ParseFun        func(r *Res, respStr string) interface{}
}

// 初始化一个默认配置
func NewConfig(url string) *Config {
	var conf = &Config{
		Method:  "GET",
		Url:     url,
		Timeout: 10 * time.Second,
		Headers: map[string]string{
			"user-agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X)",
		},
		AutoParseCookie: true,
		Data:            nil,
		Retry:           1, //默认重试一次
		ExecCount:       0,
	}
	return conf
}

// 初始化一个方法配置
func NewMethodConfig(url string, method string, data interface{}, isForm bool) *Config {
	var conf = NewConfig(url)
	conf.Method = method
	conf.Data = data
	conf.PostForm = isForm
	return conf
}

// 初始化一个Post配置
func NewPostConfig(url string, data interface{}, isForm bool) *Config {
	return NewMethodConfig(url, "POST", data, isForm)
}
