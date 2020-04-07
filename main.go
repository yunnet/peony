package main

import (
	"github.com/astaxie/beego"
	_ "github.com/yunnet/walkdog/routers"
	_ "github.com/yunnet/walkdog/sysinit"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.Log.AccessLogs = true
		beego.BConfig.Log.EnableStaticLogs = true
	}

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	//启用Session
	beego.Run()
}
