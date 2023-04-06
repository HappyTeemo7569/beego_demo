package util

import (
	"net/http"
	"strings"
)

// 获取客户端的IP
func GetClientIP(header http.Header) string {
	ip := "unknown"
	//这个提到最前面，作为优先级,nginx代理会获取到用户真实ip,发在这个环境变量上，必须要nginx配置这个环境变量HTTP_X_FORWARDED_FOR
	if header.Get("X-Forwarded-For") != "" {
		ip = header.Get("X-Forwarded-For")
	} else if header.Get("Remote_addr") != "" { //在nginx作为反向代理的架构中，使用REMOTE_ADDR拿到的将会是反向代理的的ip，即拿到是nginx服务器的ip地址。往往表现是一个内网ip。
		ip = header.Get("Remote_addr")
	} else if header.Get("Client_Ip") != "" { //HTTP_CLIENT_IP攻击者可以伪造一个这样的头部信息，导致获取的是攻击者随意设置的ip地址。
		ip = header.Get("Client_Ip")
	}

	if strings.Index(ip, ",") > 0 {
		arr := strings.Split(ip, ",")
		if len(arr) >= 1 {
			ip = arr[0]
		}
	}
	return ip
}
