package tests

import (
	"net/url"
	"testing"
	"zhaojunlike/gt-axios/crypto/tls"
	"zhaojunlike/gt-axios/net/http"
	"zhaojunlike/gt-axios/net/http/cookiejar"
	"zhaojunlike/gt-axios/net/http2"
)

func TestV3(t *testing.T) {
	var tlsConf = &tls.Config{
		NextProtos:               []string{"h2", "http/1.1"},
		SessionTicketsDisabled:   false,
		InsecureSkipVerify:       true,
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS11,
		MaxVersion:               tls.VersionTLS13,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256, tls.CurveP384},
		CipherSuites: []uint16{
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		},
	}
	//var tr = &http2.Transport{
	//	TLSClientConfig: tlsConf,
	//	DialTLS: func(network, addr string, cfg *tls2.Config) (conn net.Conn, e error) {
	//		//连接，然后做代理池
	//		return connToProxy("127.0.0.1", "8888", "www.nike.com", "443", "asdsd")
	//	},
	//}
	var ul, _ = url.Parse("http://127.0.0.1:8888")
	var px = http.ProxyURL(ul)

	var tr = &http.Transport{
		Proxy:           px,
		TLSClientConfig: tlsConf,
	}

	//把协议进行升级，那么我可以调包里面的协议
	_ = http2.ConfigureTransport(tr)

	var client = http.Client{
		Transport:     tr,
		CheckRedirect: nil,
		Timeout:       0,
	}
	var cookieJar, _ = cookiejar.New(nil)
	client.Jar = cookieJar
}
