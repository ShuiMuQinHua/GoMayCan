package routers

import (
	"GoMayCan/candivide/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/cooperation/report", &controllers.StartController{}, "*:CheckNewStartDeviceInfo")
	beego.Router("/cooperation/getTerminalInfo", &controllers.StartController{}, "*:GetTerminalInfo")
}
