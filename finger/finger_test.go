package finger

import (
	"fmt"
	"testing"
	"zhaojunlike/gt-axios/chttp"
	"zhaojunlike/gt-axios/crypto/tls"
)

func TestApi_Sll(t *testing.T) {
	var api = NewFingerApi(&tls.HelloOkHttp3_v3, chttp.LocalProxy)
	var res = api.Sll()
	fmt.Println(res)
}
