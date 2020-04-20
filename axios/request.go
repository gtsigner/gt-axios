//author: https://github.com/zhaojunlike
//date: 2019/12/11
//enable global dns cached
package axios

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"reflect"
	"strings"
	"time"
	"zhaojunlike/gt-axios/chttp"
	"zhaojunlike/gt-axios/net/http"
)

//http 请求
func SendHttpRequest(client *http.Client, config *chttp.Config) (*chttp.Res, error) {
	ts := time.Now()
	res := chttp.NewHttpRes()
	defer func() {
		te := time.Now()
		tw := te.Sub(ts)
		res.Time = int(tw.Milliseconds())
	}()
	res.Req = &chttp.Req{Config: config}
	res.Url = config.Url
	res.Message = "null"
	var req *http.Request
	var err error
	//set default method
	if config.Method == "" {
		config.Method = "GET"
	}

	//useragent
	if config.Headers == nil {
		config.Headers = map[string]string{}
	}

	//转小写
	var headers = http.Header{}
	for k, v := range config.Headers {
		headers.Set(k, v)
	}

	if config.Data == nil {
		//GET 之类的
		req, err = http.NewRequest(config.Method, config.Url, nil)
	} else {
		//包含数据的请求
		var ct = "application/json"
		//默认
		if config.PostForm {
			//判断data是否是string //TODO 判断类型，自动解析成form
			var tp = reflect.TypeOf(config.Data).String()
			var str = ""
			if tp == "string" {
				str = config.Data.(string)
			} else if tp == "url.Values" {
				str = config.Data.(url.Values).Encode()
			} else {
				panic("请求数据类型错误")
			}
			req, err = http.NewRequest(config.Method, config.Url, strings.NewReader(str))
			ct = "application/x-www-form-urlencoded; charset=UTF-8"
		} else {
			var str, err = json.Marshal(config.Data)
			if err == nil {
				req, err = http.NewRequest(config.Method, config.Url, bytes.NewReader(str))
			}
		}
		//设置请求
		if err == nil {
			headers.Set("Content-Type", ct)
		}
	}
	//整体判断错误
	if err != nil {
		res.Message = err.Error()
		return res, err
	}
	req.Header = headers
	rsp, err := client.Do(req)
	if err != nil {
		res.Message = err.Error()
		return res, nil
	}

	//close
	if rsp.Body != nil {
		defer rsp.Body.Close()
	}

	//只要不是网络出错就ok
	res.StatusCode = rsp.StatusCode
	res.Ok = rsp.StatusCode <= 540
	res.Proto = rsp.Proto
	res.Message = rsp.Status
	res.Headers = rsp.Header
	res.Success = false

	//body
	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		res.Message = fmt.Sprintf("request fail: %v", err)
		return res, nil
	}
	res.RespStr = string(buf)
	res.Buffer = buf
	ct := rsp.Header["Content-Type"]

	//TODO 优化一下逻辑，根据content-type 进行自动解析，或者通过jsonparse 强制解析
	//全部尝试用json解析
	if config.ParseFun == nil {
		var obj interface{}
		if nil == json.Unmarshal(buf, &obj) {
			res.Data = obj
		}

		//判断是否需要解析成JSON
		if config.ParseJson == false && ct != nil && len(ct) > 0 {
			cty := ct[0]
			var obj interface{}
			if strings.Index(cty, "json") != -1 && nil == json.Unmarshal(buf, &obj) {
				res.Data = obj
			}
		}
	} else {
		//调用自定义的解析方法
		res.Data = config.ParseFun(res, res.RespStr)
	}
	return res, nil
}
