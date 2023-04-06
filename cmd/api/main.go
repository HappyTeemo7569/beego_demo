package main

import (
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//run, _ := config.String("runmode")
	//if os.Getenv("runMode") != "" {
	//	run = os.Getenv("runMode")
	//}

	//core.New(core.GetFileConf(run), run)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = false
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.BConfig.RunMode = "test"
	}

	// 自定义错误 暂时不用
	//beego.BConfig.RecoverFunc = middleware.RecoverPanic
	beego.Run()
}
