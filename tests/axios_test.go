package tests

import (
	"fmt"
	"testing"
	"zhaojunlike/gt-axios/axios"
	"zhaojunlike/gt-axios/chttp"
)

func TestAxios(t *testing.T) {
	var opt = axios.NewOptions()
	opt.Proxy = chttp.LocalProxy
	var client, _ = axios.NewHttpClient(opt)
	var req = chttp.NewConfig("https://baidu.com")
	res, err := client.Request(req)
	fmt.Println(res, err)
}
