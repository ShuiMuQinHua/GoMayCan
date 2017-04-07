package model

import (
	"fmt"
	. "my_go_api/store"
)

type Aboutus struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Codetitle string `json:"codetitle"`
	Code      string `json:"QRcode"`
	Phone     string `json:"hotline"`
	Email     string `json:"serviceMail"`
	Ctime     int64  `json:"ctime"`
	Address   string `json:"address"`
}

func GetOurData() ([]*Aboutus, error) {
	aboutus := make([]*Aboutus, 0)
	err := Orm.Where("id=?", 33).Find(&aboutus)
	fmt.Println(aboutus)
	return aboutus, err
}
