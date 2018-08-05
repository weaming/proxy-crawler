package storage

import (
	"strconv"
	"time"

	"github.com/weaming/golib/proxy"
)

func GetProxyURLFromIP(ip *IP) string {
	if ip != nil {
		proxy := ip.Protocol + "://" + ip.IP + ":" + strconv.Itoa(ip.Port)
		return proxy
	}
	return ""
}

func IsValidIP(ip *IP) bool {
	proxyURL := GetProxyURLFromIP(ip)
	if proxyURL == "" {
		return false
	}

	println(proxyURL)

	for i := 0; i < 3; i++ {
		if proxy.IsValidProxy("https://www.douban.com", proxyURL, 200, 3) {
			return true
		}
		time.Sleep(500 * time.Millisecond)
	}
	return false
}
