package chttp

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	sp = "[,|，|\t|:]"
)

var (
	spReg, _           = regexp.Compile(sp)
	ErrStrNotNull      = errors.New("str can't be null")
	ErrStrLen          = errors.New("str split len must be 2 or 4")
	ErrProxyPortError  = errors.New("port error")
	LocalProxy, _      = ParseProxyStr("127.0.0.1:8888")
	ErrParseProxy      = errors.New("parse proxy fail")
	ErrNotSetCookieJar = errors.New("not set cookie jar")
)

//格式1：socks5:ip:port:username:password
//格式2：socks5:ip:port
//格式3: ip:port:username:password
//解析proxy
func ParseProxyStr(str string) (*Proxy, error) {
	var proxy = &Proxy{}
	str = strings.Trim(str, " ")
	if str == "" {
		return proxy, ErrStrNotNull
	}
	//var arr = strings.Split(str, ":")
	var arr = spReg.Split(str, -1)
	var ct = len(arr)

	if ct != 2 && ct != 4 && ct != 3 && ct != 5 {
		return proxy, ErrStrLen
	}

	var xProtocol, xHost, xPort, xUsername, xPassword int
	xProtocol = -1
	xHost = -1
	xPort = -1
	xUsername = -1
	xPassword = -1

	//1.初始2-4无protocol格式的代理
	if ct == 2 || ct == 4 {
		xHost = 0
		xPort = 1
		if ct == 4 {
			xUsername = 2
			xPassword = 3
		}
	}

	//2.有protocol格式
	if ct == 3 || ct == 5 {
		xProtocol = 0
		xHost = 1
		xPort = 2
		if ct == 5 {
			xUsername = 3
			xPassword = 4
		}
	}

	if xProtocol != -1 {
		proxy.Protocol = ProxyType(arr[xProtocol])
	} else {
		proxy.Protocol = ProxyHttp
	}
	proxy.Host = arr[xHost]
	var port, err = strconv.Atoi(arr[xPort])
	if err != nil {
		return proxy, err
	}
	proxy.Port = port
	if xUsername != -1 && xPassword != -1 {
		proxy.Username = arr[xUsername]
		proxy.Password = arr[xPassword]
		proxy.NeedAuth = true
	}

	//判断端口是否解析成功
	if proxy.Port <= 0 || proxy.Port > 65535 {
		return nil, ErrProxyPortError
	}
	return proxy, nil
}
