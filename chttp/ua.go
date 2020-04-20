package chttp

import "strings"

func FilterUserAgent(useragent string) string {
	if useragent == "" ||
		strings.Index(useragent, "spider") != -1 ||
		strings.Index(useragent, "Spider") != -1 ||
		strings.Index(useragent, "Bot") != -1 ||
		strings.Index(useragent, "bot") != -1 ||
		strings.Index(useragent, "Hexometer") != -1 ||
		strings.Index(useragent, "Macintosh") != -1 || //过滤mac os
		strings.Index(useragent, "googleweblight") != -1 {
		useragent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36"
	}
	return useragent
}
