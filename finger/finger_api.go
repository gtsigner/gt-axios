package finger

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"zhaojunlike/common"
	"zhaojunlike/common/chttp"
	"zhaojunlike/gt-axios/axios"
	tls "zhaojunlike/gt-axios/crypto/tls"
)

type Api struct {
	client *axios.Client
}

func NewFingerApi(hello *tls.ClientHelloID, proxy *chttp.Proxy) *Api {
	var opt = axios.NewOptions()
	opt.Proxy = proxy
	opt.UseHttp2 = true
	opt.TLSClientConfig = &tls.Config{
		ClientHelloID: hello,
	}
	var api = &Api{}
	var err error
	api.client, err = axios.NewHttpClient(opt)
	if err != nil {
		return nil
	}
	return api
}
func (api *Api) Ip() *chttp.Res {
	var conf = chttp.NewConfig("https://browserleaks.com/ip")
	res, _ := api.client.Request(conf)
	if res.Ok && res.RespStr != "" {
		var mapper = map[string]string{}
		//解析
		var keys = "OS,MTU,Link,IP address"
		var no = true
		doc, _ := goquery.NewDocumentFromReader(bytes.NewReader([]byte(res.RespStr)))
		doc.Find(".wb tr").Each(func(i int, selection *goquery.Selection) {
			var key = ""
			var val = ""
			selection.Find("td").Each(func(i int, sele *goquery.Selection) {
				if i == 0 {
					key = sele.Text()
				}
				if i == 1 {
					val = sele.Text()
				}
			})
			if strings.Index(keys, key) != -1 || no {
				mapper[key] = val
			}
		})
		res.Data = mapper
		res.Success = true
	}
	return res
}

func (api *Api) Sll() *chttp.Res {
	var uri = fmt.Sprintf("https://tls.browserleaks.com/json?_=%v", common.CreateTimestamp())
	var conf = chttp.NewConfig(uri)
	res, _ := api.client.Request(conf)
	if res.Ok && res.Data != nil {
		res.Success = true
	}
	return res
}
