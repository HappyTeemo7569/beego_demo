package util

import (
	"github.com/beego/beego/v2/core/logs"
	"time"
)

// ApiLog Api 追踪日志
type ApiLog struct {
	Url      string    `json:"url"`
	Header   string    `json:"header"`
	Para     string    `json:"para"`
	Context  string    `json:"context"`
	TraceId  string    `json:"trace_id"`
	Interval float64   `json:"interval"`
	Create   time.Time `json:"create"`
}

// 可以自己定义分割模式
func (l *ApiLog) ToString() string {
	return JsonStr(l)
}

func (l *ApiLog) Write() {
	logs.Notice(l.ToString())
}
