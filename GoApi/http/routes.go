package http

//github.com/gorilla/mux  用来定义路由
import (
	"my_go_api/api"
	"my_go_api/api/aboutus"
	"my_go_api/api/special"

	"github.com/gorilla/mux"
)

//把路由配置到r上
func ConfigRouter(r *mux.Router) {
	configUserRoutes(r)
}

func configUserRoutes(r *mux.Router) {
	r.HandleFunc("/user", api.GetAllUserList).Methods("GET")
	r.HandleFunc("/add", api.AddUserInfo).Methods("GET")
	r.HandleFunc("/appstore/aboutus", aboutus.GetOurInfo).Methods("GET")
	r.HandleFunc("/appstore/speciallist", special.GetSpecialList).Methods("GET")
	r.HandleFunc("/appstore/specialcontent", special.GetSpecialContent).Methods("GET")
	r.HandleFunc("/appstore/topic", topic.GetTopicList).Methods("GET")
}
