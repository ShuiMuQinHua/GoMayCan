package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type RpLog struct {
	Id              int64     `json:"id"`
	Factory         string    `json:"factory"`
	Hardwaretype    string    `json:"hardwareType"`
	Hardwareversion string    `json:"hardwareVersion"`
	Mac             string    `json:"mac"`
	Channel         string    `json:"channel"`
	Ctime           time.Time `json:"ctime"`
	Ip              string    `json:"ip"`
	Cdate           string    `json:"cdate"`
	Cpuid           string    `json:"cpuId"`
	Hdcpkey         string    `json:"hdcpKey"`
	Backmac         string    `json:"backMac"`
	Publickey       string    `json:"publicKey"`
}

func (this *ReportLog) ReportLogAdd(reportlog *ReportLog) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(reportlog)
	return id, err
}
