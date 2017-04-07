package main

import (
	_ "GoMayCan/candivide/docs"
	_ "GoMayCan/candivide/models"
	_ "GoMayCan/candivide/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
