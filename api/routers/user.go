package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"teemo_beego_demo/api/middleware"
)

func InitUserRouter() []beego.LinkNamespace {
	return []beego.LinkNamespace{
		beego.NSNamespace("/user/",
			beego.NSBefore(middleware.BeforeCheckLogin),
			userInfo(),
		),
	}
}

func userInfo() beego.LinkNamespace {
	return beego.NSNamespace("/user-info",
		//beego.NSRouter("/get", &controllers.UserController{}, "post:GetUerInfo"),
	)
}
