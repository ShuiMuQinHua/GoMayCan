package models

import (
	//	"time"

	"github.com/astaxie/beego/orm"
)

type Tndor struct {
	Vendorid int64  `orm:"pk"`
	Name     string `json:"name"`
	Ctime    int64  `json:"ctime"`
	Desc     string `json:"desc"`
}

func (this *Tndor) GetTmsVenderInfoByID(vendorId int64) []*TmsVendor {
	o := orm.NewOrm()
	tmsvendor := make([]*TmsVendor, 0)
	o.QueryTable("tms_vendor").Filter("vendorId", vendorId).All(&tmsvendor)
	return tmsvendor
}
