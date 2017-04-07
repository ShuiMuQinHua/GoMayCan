package http

import (
	"my_go_api/g"
	"my_go_api/http/render"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/urlooker/web/http/middleware"
)

//开启http服务
func Start() {
	//设置视图文件目录，后缀名，模板标签，等信息
	//赋值给了全局变量 Render
	render.Init()
	//cookie.Init()

	//涉及到路由 返回一个Router结构体
	r := mux.NewRouter().StrictSlash(false)
	//把设置好的路由信息，绑定到r上
	ConfigRouter(r)

	//返回一个Negroni结构体对象
	n := negroni.New()
	//middleware.NewRecovery() 返回一个Recovery结构体的对象
	n.Use(middleware.NewRecovery())
	//添加路由信息
	n.UseHandler(r)
	//启动0.0.0.0:1984端口
	n.Run(g.Config.Http.Listen)
}
