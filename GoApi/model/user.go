package model

import (
	"fmt"

	. "my_go_api/store"
)

type User struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
	Roles    string `json:"roles"`
}

var UserItem *User

func Insert() int64 {
	user := User{UserName: "cyy", Pwd: "1234", Roles: "123"}
	var aa int64
	for i := 0; i < 10000000; i++ {
		aa, _ = Orm.Insert(&user)
		fmt.Println(aa)
	}
	return aa
}

func GetUserById(userId int64) ([]*User, error) {
	fmt.Println(userId)
	users := make([]*User, 0)
	err := Orm.Where("id = ?", userId).Find(&users)
	return users, err
}
