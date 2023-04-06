package middleware

import (
	"github.com/beego/beego/v2/server/web/context"
)

type Header struct {
	PhoneBrand     string `json:"phoneBrand"`     //手机品牌
	PhoneSystem    string `json:"phoneSystem"`    //操作系统
	PhoneModels    string `json:"phoneModels"`    //手机型号
	AppMarket      string `json:"appMarket"`      //渠道
	AppVersionCode string `json:"appVersionCode"` //app版本号
	AppVersionName string `json:"appVersionName"` //app版本名字
	ApiVersion     string `json:"apiVersion"`     //api版本
	PackageName    string `json:"packageName"`    //包名

	AppType     string `json:"appType"`
	DeviceType  string `json:"deviceType"`
	LocalIp     string `json:"localIp"`
	IsTablet    string `json:"isTablet"`
	Isp         string `json:"isp"`
	Uuid        string `json:"uuid"`
	NetworkType string `json:"networkType"`
	SysVersion  string `json:"sysVersion"`
	AppId       string `json:"appId"`
}

// CheckHeader 头部校验
var CheckHeader = func(ctx *context.Context) {
	//us := NewUserContext(ctx)
	//us.save()
}
