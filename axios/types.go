//author: https://github.com/zhaojunlike
//date: 2019/12/11
package axios

import (
    "time"
    "zhaojunlike/common/chttp"
    "zhaojunlike/gt-axios/crypto/tls"
    "zhaojunlike/gt-axios/net/http/cookiejar"
)

type Options struct {
    Proxy           *chttp.Proxy
    Timeout         time.Duration
    AllowRedirect   bool
    CookieJar       *cookiejar.Jar
    tslConfig       *tls.Config
    UseHttp2        bool
    Debug           bool
    TLSClientConfig *tls.Config
    DnsCache        bool //是否开启Dns缓存
}

func NewOptions() *Options {
    return &Options{Proxy: nil, Timeout: 30 * time.Second, AllowRedirect: false, UseHttp2: true, DnsCache: true}
}
