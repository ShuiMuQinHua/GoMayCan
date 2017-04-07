package api

import (
	"fmt"
	"net/http"
	"strconv"

	"my_go_api/http/params"
	"my_go_api/http/render"
	"my_go_api/model"
)

func GetAllUserList(w http.ResponseWriter, r *http.Request) {

	id := params.MustString(r, "id")
	// name := params.MustString(r, "name")
	idint, _ := strconv.ParseInt(id, 10, 64)
	userList, _ := model.GetUserById(idint)
	fmt.Println(userList)
	// data := map[string]string{"id": id, "name": name}

	render.JSON(w, userList, 200)
}

func AddUserInfo(w http.ResponseWriter, r *http.Request) {

	a := model.Insert()
	fmt.Println(a)
}

//func mulitAddUserInfo(w http.ResponseWriter, r *http.Request) {
//	for i = 0; i < 50; i++ {
//		go AddUserInfo()
//	}
//}
