package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"teemo_beego_demo/api/middleware"
)

func InitRouter() *beego.Namespace {
	r := make([]beego.LinkNamespace, 0)

	r = append(r,
		beego.NSBefore(middleware.CheckHeader),
	)

	r = append(r, InitUserRouter()...)

	//可以添加统一前缀
	return beego.NewNamespace("/demo",
		r...,
	)
}
