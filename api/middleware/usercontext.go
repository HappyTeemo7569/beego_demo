package middleware

import (
	"errors"
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"github.com/HappyTeemo7569/teemoKit/util"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"reflect"
	"time"
)

type UserContext struct {
	Context
	Header Header
	UserId int64
}

func (this *UserContext) save() {
	this.Ctx.Input.Data()["_bcontext"] = *this
}

func (this *UserContext) logWrite(c interface{}) {

	var param string
	if this.Header.ApiVersion == "web" {
		param = util.JsonStr(this.Ctx.Input.RequestBody) //web用json
	} else {
		param = util.JsonStr(this.Ctx.Request.Form) //客户端用from表单请求
	}

	api := tlog.ApiLog{
		this.Ctx.Request.RequestURI,
		util.JsonStr(this.Header),
		param,
		util.JsonStr(c),
		this.TraceId,
		float64(time.Now().Sub(this.CreateTime).Milliseconds()),
		this.CreateTime,
	}
	if api.Context == "\"请求结束：\"" {
		api.Context = api.Context + util.JsonStr(this.Ctx.Input.Data()["json"])
	}

	api.Write()
}

// 结束输出
func (this *UserContext) Json(data interface{}) {
	this.Ctx.Input.Data()["json"] = data
	err := this.Ctx.Output.JSON(data, true, false)
	if err != nil {
		return
	}
}

func NewUserContext(ctx *context.Context) *UserContext {
	rq := UserContext{
		NewContext(ctx),
		GetHead(ctx.Request.Header),
		0,
	}
	rq.BOutput = &rq
	return &rq
}

// 获取头
func GetHead(hd http.Header) Header {
	h := new(Header)
	valueElem := reflect.ValueOf(h).Elem()

	typeElem := valueElem.Type()
	for i := 0; i < valueElem.NumField(); i++ {
		name := typeElem.Field(i).Tag.Get("json") //将tag输出出来
		valueElem.Field(i).SetString(hd.Get(name))
	}
	return *h
}

// 获取content
func LoadContext(ctx *context.Context) UserContext {
	if rq, ok := ctx.Input.Data()["_bcontext"].(UserContext); ok {
		return rq
	}
	panic(errors.New("获取Context出错！"))
}
