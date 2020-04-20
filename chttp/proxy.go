package chttp

import "fmt"

type Proxy struct {
	Host     string    `json:"host"`
	Port     int       `json:"port"`
	Username string    `json:"username,omitempty"`
	Password string    `json:"password,omitempty"`
	NeedAuth bool      `json:"need_auth,omitempty"`
	Protocol ProxyType `json:"protocol,omitempty"` //协议
}

func (p Proxy) String() string {
	if p.NeedAuth {
		return fmt.Sprintf("%v:%v:%v:%v:%v", p.Protocol, p.Host, p.Port, p.Username, p.Password)
	}
	return fmt.Sprintf("%v:%v:%v", p.Protocol, p.Host, p.Port)
}

type ProxyType string

var (
	ProxySocks5 ProxyType = "socks5"
	ProxySocks4 ProxyType = "socks4"
	ProxyHttps  ProxyType = "https"
	ProxyHttp   ProxyType = "http"
)

func GetProxyTypes() []ProxyType {
	var types = []ProxyType{ProxySocks5, ProxySocks4, ProxyHttp, ProxyHttps}
	return types
}
