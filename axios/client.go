package axios

import (
	"context"
	"fmt"
	"github.com/viki-org/dnscache"
	"net"
	"net/url"
	"strings"
	"sync"
	"time"
	"zhaojunlike/gt-axios/chttp"
	tls "zhaojunlike/gt-axios/crypto/tls"
	"zhaojunlike/gt-axios/net/http"
	"zhaojunlike/gt-axios/net/http/cookiejar"
	"zhaojunlike/gt-axios/net/http2"
)

type Client struct {
	Http    *http.Client
	Options *Options
	locker  *sync.Mutex
}

var resolver = dnscache.New(time.Minute * 5)

//构造HttpClient
func NewHttpClient(opt *Options) (*Client, error) {
	var tsConf = opt.TLSClientConfig
	if tsConf == nil {
		tsConf = &tls.Config{}
	}
	if tsConf.ClientHelloID == nil {
		tsConf.ClientHelloID = &tls.HelloGolang
	}

	tr := &http.Transport{
		TLSClientConfig:     tsConf,
		MaxIdleConnsPerHost: 128,
		TLSHandshakeTimeout: 30 * time.Second,
		ProxyConnectHeader: map[string][]string{
			"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36"},
		},
	}
	if opt.DnsCache {
		tr.DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
			separator := strings.LastIndex(addr, ":")
			ip, _ := resolver.FetchOneString(addr[:separator])
			return net.Dial("tcp", ip+addr[separator:])
		}
	}

	//有些默认直接干掉
	if opt.Proxy != nil && opt.Proxy.Port <= 0 {
		opt.Proxy = nil
	}

	//使用代理
	if opt.Proxy != nil {
		uri := url.URL{}
		var proxyUrl *url.URL
		var urlStr = "http://"

		//默认使用http代理
		if opt.Proxy.Protocol == "" {
			opt.Proxy.Protocol = chttp.ProxyHttp
		}
		urlStr = string(opt.Proxy.Protocol) + "://"
		if opt.Proxy.Username != "" && opt.Proxy.Password != "" {
			opt.Proxy.NeedAuth = true
			urlStr += fmt.Sprintf("%s:%s@", opt.Proxy.Username, opt.Proxy.Password)
		}
		urlStr += fmt.Sprintf("%s:%d", opt.Proxy.Host, opt.Proxy.Port)
		proxyUrl, _ = uri.Parse(urlStr)

		if proxyUrl != nil {
			tr.Proxy = http.ProxyURL(proxyUrl)
		}
	}

	//使用http2
	if opt.UseHttp2 {
		if err := http2.ConfigureTransport(tr); err != nil {
			return nil, err
		}
	}

	//超时
	if opt.Timeout <= 0 {
		opt.Timeout = 20 * time.Second
	}

	client := Client{Options: opt, locker: new(sync.Mutex)}
	client.Http = &http.Client{
		Transport: tr,
		Timeout:   opt.Timeout,
	}
	if opt.CookieJar != nil {
		client.Http.Jar = opt.CookieJar
	}

	//禁止自动转发
	if opt.AllowRedirect == false {
		client.Http.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	return &client, nil
}

func (client *Client) Destroy() {
	client.Http.CloseIdleConnections()
}

func (client *Client) Close() {
	client.Http.CloseIdleConnections()
}

//发送请求
func (client *Client) Request(config *chttp.Config) (res *chttp.Res, err error) {
	//request
	for {
		res, err = SendHttpRequest(client.Http, config)
		config.ExecCount++
		if res == nil {
			break
		}
		//是否已经请求成功或者超出次数了
		if config.Retry < config.ExecCount || res.StatusCode < 550 {
			break
		}
	}
	return res, err
}

//set cookirjar
func (client *Client) SetCookieJar(jar *cookiejar.Jar) {
	client.Options.CookieJar = jar
	client.Http.Jar = client.Options.CookieJar
}

// 设置单个cookiejar的cookie数据
func (client *Client) SetCookie(uri string, cookie *http.Cookie) error {
	var r, err = url.Parse(uri)
	if err != nil {
		return err
	}
	var arr = make([]*http.Cookie, 1)
	arr[0] = cookie
	if client.Http.Jar != nil {
		client.Http.Jar.SetCookies(r, arr)
		return nil
	}
	return chttp.ErrNotSetCookieJar
}

// 移除COokie
func (client *Client) RemoveCookie(uri string, name string) error {
	c := &http.Cookie{Name: name, Value: "", Expires: time.Unix(0, 0)}
	return client.SetCookie(uri, c)
}

// 移除C0okie
func (client *Client) GetCookie(uri string, name string) *http.Cookie {
	var urs, _ = url.Parse(uri) //
	var cookies = client.Http.Jar.Cookies(urs)
	for _, v := range cookies {
		if v.Name == name {
			return v
		}
	}
	return nil
}

//
func (client *Client) GetCookies(uri string) []*http.Cookie {
	var urs, _ = url.Parse(uri) //
	var cookies = client.Http.Jar.Cookies(urs)
	return cookies
}

func (client *Client) SetCookies(uri string, cookies []*http.Cookie) {
	var urs, _ = url.Parse(uri) //
	client.Http.Jar.SetCookies(urs, cookies)
}
