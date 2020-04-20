//author: https://github.com/zhaojunlike
//date: 2019/12/12
package axios

import (
	"fmt"
	"testing"
	"zhaojunlike/gt-axios/chttp"
	tls "zhaojunlike/gt-axios/crypto/tls"
)

func request() {
	var opt = NewOptions()
	opt.Proxy = chttp.LocalProxy
	opt.Debug = true
	var stat = &tls.ClientSessionState{}
	opt.TLSClientConfig = &tls.Config{
		ClientSessionState: stat,
		//ClientHelloID:      &tls.HelloFirefox_65,
		ClientHelloID: &tls.HelloIOSChrome_13_3,
	}
	client, err := NewHttpClient(opt)
	fmt.Println(err)

	conf := chttp.NewConfig("https://www.nike.com")
	var headers = map[string]string{}
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/538.36"
	headers["sec-fetch-site"] = "same-origin"
	headers["sec-fetch-mode"] = "navigate"
	headers["upgrade-insecure-requests"] = "1"
	conf.Headers = headers

	res, err := client.Request(conf)
	fmt.Println(res.StatusCode)
	fmt.Println(stat)
}

func TestNewDefaultHttpRes(t *testing.T) {
	//px.Host = "223.241.88.141"
	//px.Port = 18756
	//px.Username = "zhaojunlike"
	//px.Password = "zhaojunlike"
	request()
}
