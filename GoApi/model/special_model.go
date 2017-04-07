package model

import (
	"fmt"
	. "my_go_api/store"
)

//注意字段名  除了 首字母外  其他的不能大写
type Special struct {
	Id               int64  `json:"id"`
	Spname           string `json:"spname"`
	Desc             string `json:"desc"`
	Backimage        string `json:"backImage"`
	Listimage        string `json:"ListImage"`
	Transverseimage  string `json:"TransverseImage"`
	Typeid           int64  `json:"typeId"`
	Sort             int64  `json:"Sort"`
	Status           int64  `json:"Status"`
	Ctime            int64  `json:"Ctime"`
	Utime            int64  `json:"Utime"`
	Dimensiontype    int64  `json:"DimensionType"`
	Actiontype       int64  `json:"ActionType"`
	Layoutactionid   int64  `json:"LayoutActionId"`
	Contentid        int64  `json:"ContentId"`
	Actionpara       string `json:"ActionPara"`
	Actionoutpackage string `json:"ActionOutPackage"`
	Actioncode       string `json:"ActionCode"`
}

//会自动转换为表 special_content 驼峰法的字段也会做同样的转换如SpecialId 对应转换数据库字段为special_id
type SpecialContent struct {
	Id        int64 `json:"id"`
	Specialid int64 `json:"specialId"`
	Appid     int64 `json:"appId"`
	Sort      int64 `json:"sort"`
	Ctime     int64 `json:"ctime"`
	Status    int64 `json:"status"`
	Isdel     int64 `json:"isDel"`
}

func GetSpecialList(typeID int64) ([]*Special, error) {
	fmt.Println(typeID)
	special := make([]*Special, 0)
	err := Orm.Where("typeId=?", typeID).Limit(5, 0).Find(&special)

	fmt.Println(len(special))
	return special, err
}

func GetSpecialContent(specialID int64) (interface{}, error) {
	special := new(Special)
	has, err := Orm.Where("id=?", specialID).Get(special)
	specialContent := make([]*SpecialContent, 0)
	appids := make([]int64, 0)
	appdatas := make([]*Application, 0)
	result := make(map[string]interface{})
	if has {
		fmt.Println("yes")
		err = Orm.Where("specialId=?", special.Id).Find(&specialContent)
		total := len(specialContent)
		for i := 0; i < total; i++ {
			appids = append(appids, specialContent[i].Appid)
		}
		appdatas, err = getAppDataByID(appids)
		data := make(map[string]interface{})
		data["id"] = special.Id
		data["name"] = special.Spname
		data["type"] = special.Typeid
		data["content"] = appdatas
		result["status"] = 0
		result["message"] = "success"
		result["data"] = data
	}
	return result, err
}
