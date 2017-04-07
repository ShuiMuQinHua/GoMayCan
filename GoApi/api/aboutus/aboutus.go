package aboutus

import (
	"fmt"
	"net/http"

	"my_go_api/http/render"
	"my_go_api/model"
)

func GetOurInfo(w http.ResponseWriter, r *http.Request) {
	ourinfo, _ := model.GetOurData()
	fmt.Println(ourinfo)
	res := new(model.ResponseFormat)
	if len(ourinfo) < 1 {
		data := make(map[string]string)
		res.Status = "0"
		res.Message = "fail"
		res.Data = data
	} else {
		res.Status = "1"
		res.Message = "success"
		data1 := make(map[string]string)
		data1["serviceMail"] = ourinfo[0].Email
		data1["hotline"] = ourinfo[0].Phone
		data1["QRcode"] = ourinfo[0].Code
		res.Data = data1
	}
	render.JSON(w, res, 200)
}
