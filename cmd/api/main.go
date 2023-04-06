package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"teemo_beego_demo/api/routers"
)

func main() {
	beego.AddNamespace(routers.InitRouter())

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = false
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.BConfig.RunMode = "test"
	}

	// 自定义错误 暂时不用
	//beego.BConfig.RecoverFunc = middleware.RecoverPanic
	beego.Run()
}
