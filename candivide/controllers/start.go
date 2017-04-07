package controllers

import (
	"GoMayCan/candivide/helpers"
	"GoMayCan/candivide/models"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type StartController struct {
	BaseController
}

func (this *StartController) CheckNewStartDeviceInfo() {
	ip := this.Ctx.Input.IP()
	mac := this.GetString("mac")
	data := this.GetString("data")
	fmt.Println("ip---" + ip)
	fmt.Println("mac---" + mac)
	fmt.Println("crypted---" + data)

	key := helpers.GetMyConfig("cryKey")
	fmt.Println("cryptKey" + key)
	decryptData, _ := helpers.TripleDesDecryptString(data, key)

	dataSli := strings.Split(string(decryptData[:]), "@")
	fmt.Println("dataLen" + strconv.Itoa(len(dataSli)))
	checkMac := this.CheckMac(mac)
	dataSlice := make([]interface{}, 0)
	if checkMac != 1 {
		if checkMac == -1 {
			this.RespMsg(dataSlice, ERROR_MAC_FORMAT, "mac格式错误")
		} else {
			this.RespMsg(dataSlice, ERROR_MAC_OWN, "mac为自有mac")
		}
		this.RespMsg(dataSlice, ERROR_CONFLICT, "Conflict")
	} else {
		report := new(models.ReportLog)
		t := time.Now()

		report.Factory = dataSli[1]
		report.Hardwaretype = dataSli[2]
		report.Hardwareversion = dataSli[3]
		report.Mac = mac
		report.Channel = dataSli[4]
		report.Ctime = t
		report.Ip = ip
		report.Cdate = t.Format("2006-01-02")
		report.Cpuid = dataSli[5]
		report.Hdcpkey = dataSli[6]
		report.Backmac = dataSli[7]
		report.Publickey = dataSli[8]

		addReportId, _ := report.ReportLogAdd(report)
		this.RespMsg(addReportId, ERROR_OK, "success")
	}
}

func (this *StartController) GetTerminalInfo() {
	mac := this.GetString("mac")
	data := this.GetString("data")
	key := helpers.GetMyConfig("cryKey")
	fmt.Println("mac---" + mac)
	fmt.Println("data---" + data)
	fmt.Println("key---" + key)

	decryptData, _ := helpers.TripleDesDecryptString(data, key)
	dataSli := strings.Split(string(decryptData[:]), "$")
	if mac != dataSli[0] {
		this.RespMsg("", ERROR_MAC_NOTMATCH, "mac不匹配")
	}

	terminl := new(models.Terminal)
	res, flag := terminl.GetTerminalInfo(mac)
	dataRes := make(map[string]interface{})
	dataRes["mac"] = mac
	if flag == 1 {
		dataRes["inMacPool"] = 1
		venderName := GetVenderNameById(res[0].Vendorid)
		dataRes["vendor"] = venderName
		dataRes["hardwareType"] = res[0].Hardwaretype
		firstTime := time.Unix(res[0].Vendorfirsttime, 0)
		dataRes["vendorFirstTime"] = firstTime.Format("2006-01-02 15:04:02")
		if res[0].Uservalidtime > 0 {
			dataRes["userValidStatus"] = 1
		} else {
			dataRes["userValidStatus"] = 0
		}
		this.RespMsg(dataRes, ERROR_OK, "success")
	} else {
		dataRes["inMacPool"] = 0
		dataRes["vendor"] = ""
		dataRes["hardwareType"] = ""
		dataRes["vendorFirstTime"] = ""
		dataRes["userValidStatus"] = ""
		this.RespMsg(dataRes, ERROR_NODATA, "nodata")
	}

}

func GetVenderNameById(venderID int64) string {
	tmsVendorM := new(models.TmsVendor)
	tmsVendorData := tmsVendorM.GetTmsVenderInfoByID(venderID)
	if tmsVendorData[0].Name != "" {
		return tmsVendorData[0].Name
	} else {
		return "未知"
	}
}
