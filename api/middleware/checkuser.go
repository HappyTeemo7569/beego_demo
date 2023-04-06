package middleware

import (
	"github.com/beego/beego/v2/server/web/context"
)

// 检查快速登录效验的效验器
var BeforeCheckLogin = func(ctx *context.Context) {
	rq := LoadContext(ctx)
	rq.checkLogin()
}

func (this *UserContext) checkLogin() {

	type User struct {
		UserId int64  `json:"userId"`
		Token  string `json:"token"`
	}

	//添加鉴权逻辑

	this.save()
}
