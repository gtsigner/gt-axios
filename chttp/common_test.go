package chttp

import (
	"fmt"
	"testing"
)

func TestParseProxyStr(t *testing.T) {
	var str1 = "192.168.10.66:8888:godtoy:pwd"
	px, _ := ParseProxyStr(str1)
	fmt.Println(px)
	str1 = "socks5:192.168.10.66:8888:godtoy:pwd"
	px, _ = ParseProxyStr(str1)
	fmt.Println(px)
	str1 = "http:192.168.10.66:8888"
	px, _ = ParseProxyStr(str1)
	fmt.Println(px)
	str1 = "192.168.10.66:8888"
	px, _ = ParseProxyStr(str1)
	fmt.Println(px)
}

func TestLocalProxy(t *testing.T) {
	fmt.Println(LocalProxy)
}
