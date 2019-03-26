package main

import (
	_ "github.com/yunnet/peony/routers"
	_ "github.com/yunnet/peony/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	//if beego.BConfig.RunMode == "dev" {
	//
	//}

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	//启用Session
	beego.Run()
}

