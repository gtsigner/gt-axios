## 使用utls 构造一套完整的http客户端
clone form golang src sdk 1.13.7

-   支持http2
-   支持代理
-   支持utls指纹
-   支持utls自定义指纹
-   修改http2 header头排序

## 使用方法
```shell script

go get github.com/zhaojunlike/go-http-utls

```

axios 基础包装

```golang

func TestAxios(t *testing.T) {
	var opt = axios.NewOptions()
	opt.Proxy = &axios.Proxy{Host: "127.0.0.1", Port: 8888}
	var client, _ = axios.NewHttpClient(opt)
	var req = axios.NewConfig("https://baidu.com")
	res, err := client.Request(req)
	fmt.Println(res, err)
}
```

fake tls fingerprint
```golang
func TestAxios(t *testing.T) {
	var opt = axios.NewOptions()
	opt.Proxy = &axios.Proxy{Host: "127.0.0.1", Port: 8888}
	var client, _ = axios.NewHttpClient(opt)
	var req = axios.NewConfig("https://baidu.com")
    // 使用假指纹
	req.HelloId = tls.HelloChrome_72
	res, err := client.Request(req)
	fmt.Println(res, err)
} 
```


## feature
-   request frame like tor browser
-   custom tls fingerprint
-   custom transport
