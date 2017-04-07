package models

import (
	//	"time"

	"github.com/astaxie/beego/orm"
)

type Ter struct {
	Id                 int64
	Mac                string
	Vendorid           int64
	Channelid          int64
	Externalmodelid    int64
	Internalmodelid    int64
	Equipmenttype      int64
	Hardwaretype       string
	Hardwareversion    string
	Cpuid              string
	Hdcpkey            string
	Backmac            string
	Publickey          string
	Vendorfirsttime    int64
	Vendorfirstdate    string
	Vendorvalidtime    int64
	Vendorvaliddate    string
	Vendorip           string
	Userfirsttime      int64
	Userfirstdate      string
	Uservalidtime      int64
	Uservaliddate      string
	Userip             string
	Uservaliddaynum    int64
	Validuserstatus    int64
	Validuserdate      string
	Validuserendstatus int64
	Validuserenddate   string
	Validuserpaystatus int64
	Validuserpaydate   string
	Userlasttime       int64
	Userlastip         string
	Ctime              int64
	Remarks            string
}

func (this *Terminal) GetTerminalInfo(mac string) ([]*Terminal, int64) {
	o := orm.NewOrm()
	terminal := make([]*Terminal, 0)
	o.QueryTable(TableName("terminal")).Filter("mac", mac).All(&terminal)

	if len(terminal) > 0 {
		return terminal, 1
	} else {
		return terminal, 2
	}
}
