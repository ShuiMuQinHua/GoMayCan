package helpers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/url"
	"sort"
	"strings"

	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

func MyHttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		ret := map[string]string{"code": "500", "msg": "get error occur"}
		ret_json, _ := json.Marshal(ret)
		return string(ret_json)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ret := map[string]string{"code": "500", "msg": "body error occur"}
		ret_json, _ := json.Marshal(ret)
		return string(ret_json)
	}

	return string(body)
}

func MyHttpPost(url string, param interface{}) string {
	query := ""
	switch param.(type) {
	case string:
		query = param.(string)
	default:
		query = GetUrlQuery(param.(map[string]string))
	}

	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(query))
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body error occur")
	}

	fmt.Println(string(body))
	return string(body)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetUrlQuery(params map[string]string) string {
	query := ""
	for key, val := range params {
		query += key + "=" + val + "&"
	}

	return query[0 : len(query)-1]

}

func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func SiteUrl(url string) string {
	return GetMyConfig("siteUrl") + ":" + GetMyConfig("httpport") + "/" + url
}

func GetMyConfig(key string) string {
	return beego.AppConfig.String(key)
}

func RecordLog(format string, v ...interface{}) {
	log := logs.NewLogger(1e3)
	log.SetLogger("file", `{"filename":"./run.log"}`)
	log.Async()
	log.EnableFuncCallDepth(true)

	log.Info(format, v)
}

func GetSortIndex(mapData map[string]int64) []string {
	sortIndex := make([]string, len(mapData))
	var masterVal, j int64 = 0, 0
	count := len(mapData)

	for i := 0; i < count; i++ {
		masterVal = 0
		for i2, v2 := range mapData {
			if masterVal <= v2 {
				sortIndex[j] = i2
				masterVal = v2
			}
		}
		delete(mapData, sortIndex[j])
		j++
	}

	fmt.Println(sortIndex)
	return sortIndex
}
