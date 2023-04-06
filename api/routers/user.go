package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"teemo_beego_demo/api/controllers"
	"teemo_beego_demo/api/middleware"
)

func initUserRouter() []beego.LinkNamespace {
	return []beego.LinkNamespace{
		beego.NSNamespace("/user/",
			beego.NSBefore(middleware.BeforeCheckLogin),
			userInfo(),
			userAction(),
		),
	}
}

func userInfo() beego.LinkNamespace {
	return beego.NSNamespace("/user-info",
		beego.NSRouter("/get", &controllers.UserController{}, "post:GetUserInfo"),
	)
}

func userAction() beego.LinkNamespace {
	return beego.NSNamespace("/user-action",
		beego.NSRouter("/focus", &controllers.UserController{}, "post:Focus"),
	)

}
