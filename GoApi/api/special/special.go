package special

import (
	"fmt"
	"my_go_api/http/params"
	"my_go_api/http/render"
	"my_go_api/model"
	"net/http"
	"strconv"
)

func GetSpecialList(w http.ResponseWriter, r *http.Request) {
	typeid := params.MustString(r, "typeid")
	fmt.Println(typeid)
	typeidint, _ := strconv.ParseInt(typeid, 10, 64)
	speciallist, err := model.GetSpecialList(typeidint)

	resmap := make(map[string]interface{})
	if err != nil {
		resmap["status"] = "0"
		resmap["message"] = "fail"
		resmap["data"] = make([]interface{}, 0)
	} else {
		resmap["status"] = "1"
		resmap["message"] = "success"
		resmap["data"] = speciallist
	}
	render.JSON(w, resmap, 200)
}

func GetSpecialContent(w http.ResponseWriter, r *http.Request) {
	specialid := params.MustString(r, "specialid")
	fmt.Println(specialid)
	specialidint, _ := strconv.ParseInt(specialid, 10, 64)
	speCon, _ := model.GetSpecialContent(specialidint)
	render.JSON(w, speCon, 200)
}
