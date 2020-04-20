## 类似axios api 对golang http2库 实现

支持Proxy
``` golang
func TestNewDefaultHttpRes(t *testing.T) {
	var opt = &http.Options{Proxy: &http.Proxy{Host: "127.0.0.1", Port: 8888}}
	client, _ := http.NewHttpClient(opt)
	defer client.Destroy()
	ids := []string{"d8469d31-ca22-474b-a329-450d32adc789", "d8469d31-ca22-474b-a329-450d32adc789", "d8469d31-ca22-474b-a329-450d32adc789"}
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		uri := fmt.Sprintf("https://baidu.com/%s", id)
		conf := http.NewConfig(uri)
		res, _ := client.Request(conf)
		res.Println()
		wg.Done()
	}
	wg.Wait()
	fmt.Println("ressss")
}
```

-   axios 风格api
-   支持http2
-   支持proxy
-   支持retry重试

TODO 适配http1.1 1 proxy




## common

``` text
# 解析proxy字符串
func ParseProxy(str string) *Proxy
```
