package middleware

import (
	"github.com/beego/beego/v2/server/web/context"
	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
	"teemo_beego_demo/internal/core"
	"teemo_beego_demo/internal/util"
	"time"
)

// 通用输出接口
type BOutput interface {
	Json(data interface{})
}

type Context struct {
	Ctx        *context.Context
	ClientIp   string //客户端IP
	CreateTime time.Time
	TraceId    string //链路ID
	BOutput
}

func NewContext(ctx *context.Context) Context {
	trace, _ := uuid.NewV4()
	rq := Context{
		Ctx:        ctx,
		ClientIp:   util.GetClientIP(ctx.Request.Header),
		CreateTime: time.Now(),
		TraceId:    trace.String(),
	}
	return rq
}

type ResultData struct {
	Result bool        `json:"result"`
	Data   interface{} `json:"data,omitempty"`
}

func (this *Context) DataOrError(data interface{}, e *core.Error) {
	if e != nil {
		this.ThrowError(e)
		return
	}
	this.BOutput.Json(ResultData{true, data})
}

func (this *Context) TrueOrError(data bool, e *core.Error) {
	if e != nil {
		this.ThrowError(e)
		return
	}
	this.BOutput.Json(ResultData{Result: data})
}

func (this *Context) TrueBase() {
	this.BOutput.Json(ResultData{Result: true})
}
func (this *Context) TrueMessage(m string) {
	this.BOutput.Json(ResultMessage{true, m})
}

func (this *Context) TrueMessageData(m string, data interface{}) {
	this.BOutput.Json(ResultMessageData{true, m, data})
}

type ResultMessageData struct {
	Result bool        `json:"result"`
	Msg    string      `json:"txtMessage"`
	Data   interface{} `json:"data,omitempty"`
}

type ResultMessage struct {
	Result bool   `json:"result"`
	Msg    string `json:"txtMessage"`
}

type ResultError struct {
	Result bool `json:"result"`
	*core.Error
}

func (this *Context) ThrowError(e *core.Error) {
	this.BOutput.Json(ResultError{false, e})
}

func (this *Context) ThrowErrorMessage(m string) {
	this.BOutput.Json(ResultError{false, core.NewError(m)})
}

func (this *Context) ThrowErrorCode(m string, c core.ErrorCode) {
	this.BOutput.Json(ResultError{false, core.NewError(m).AppendCode(c)})
}

func (this *Context) Page(data interface{}, currPage int, isEnd bool) {
	type PagingData struct {
		Result   bool        `json:"result"`
		CurrPage int         `json:"currPage"` //当前页
		IsEnd    bool        `json:"isEnd"`    //是否数据结束
		Data     interface{} `json:"data"`     //列表数据数组
	}
	this.BOutput.Json(PagingData{
		true,
		currPage,
		isEnd,
		data,
	})
}
