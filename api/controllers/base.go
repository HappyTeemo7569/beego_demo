package controllers

import (
	"encoding/json"
	"errors"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/go-playground/validator/v10"
	"teemo_beego_demo/api/middleware"
	"teemo_beego_demo/internal/core"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) LoadContext() middleware.UserContext {
	if rq, ok := this.Ctx.Input.Data()["_bcontext"].(middleware.UserContext); ok {
		return rq
	}
	panic(errors.New("获取Context对象出错！"))
}

func (this *BaseController) CheckFormAndGet(d interface{}) bool {
	rq := this.LoadContext()

	if err := this.ParseForm(d); err != nil {
		rq.ThrowError(core.NewError("参数错误!"))
		logs.Error("参数错误!", this.Ctx.Request.Form, err)
		return false
	}

	validate := validator.New()
	err := validate.Struct(d)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			rq.ThrowError(core.NewError("参数效验错误!"))
			return false
		}
	}

	return true
}

func (this *BaseController) CheckJsonAndGet(d interface{}) bool {
	rq := this.LoadContext()

	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &d)
	if err != nil {
		rq.TrueMessage("json.Unmarshal is err:" + err.Error())
		return false
	}

	validate := validator.New()
	err = validate.Struct(d)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			rq.ThrowError(core.NewError("参数效验错误!"))
			return false
		}
	}

	return true
}
