package controllers

import (
	"fmt"
	"regexp"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type Response struct {
	Data      interface{} `json:"data"`
	Error     string      `json:"Message"`
	ErrorCode int         `json:"Status"`
}

func (this *BaseController) JsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) RespMsg(data interface{}, errorCode int, errorMsg ...string) {
	var resInfo string
	if len(errorMsg) > 0 {
		resInfo = errorMsg[0]
	}

	res := Response{
		Data:      data,
		ErrorCode: errorCode,
		Error:     resInfo,
	}
	this.JsonResult(res)
}

func (this *BaseController) CheckMac(mac string) int64 {
	return 1
	pattern := ""
	isMatch, _ := regexp.MatchString(pattern, mac)
	fmt.Println(isMatch)
	if isMatch {
		return 1
	} else {
		return -1
	}
}
