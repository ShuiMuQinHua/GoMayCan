package model

import (
	//	"fmt"
	. "my_go_api/store"
)

type Application struct {
	Id              int64  `json:"id"`
	Title           string `json:"title"`
	Appname         string `json:"Appname"`
	Appdesc         string `json:"Appdesc"`
	Appdetail       string `json:"Appdetail"`
	Recommend       string `json:"Recommend"`
	Filesize        int64  `json:"Filesize"`
	Versionname     int64  `json:"Versionname"`
	Developer       int64  `json:"Developer"`
	Devicename      int64  `json:"Devicename"`
	Listicon        int64  `json:"Listicon"`
	Detailicon      int64  `json:"Detailicon"`
	Appimage        int64  `json:"Appimage"`
	Pakagename      int64  `json:"Pakagename"`
	Pakageurl       int64  `json:"Pakageurl"`
	Downloadurl     string `json:"Downloadurl"`
	Basedownloadcnt string `json:"Basedownloadcnt"`
	Screenshot1     string `json:"Screenshot1"`
	Screenshot2     int64  `json:"Screenshot2"`
	Screenshot3     int64  `json:"Screenshot3"`
	Screenshot4     int64  `json:"Screenshot4"`
	Screenshot5     int64  `json:"Screenshot5"`
	Status          int64  `json:"Status"`
	Ctime           int64  `json:"Ctime"`
	Initial         int64  `json:"Initial"`
	Cornid          string `json:"Cornid"`
	Commentlevel    string `json:"Commentlevel"`
	Changetype      string `json:"Changetype"`
	Appmd5          int64  `json:"Appmd5"`
	Versioncode     int64  `json:"Versioncode"`
	Updatedesc      string `json:"Updatedesc"`
}

func getAppDataByID(ids []int64) ([]*Application, error) {
	application := make([]*Application, 0)
	err := Orm.In("id", ids).Find(&application)
	return application, err
}
