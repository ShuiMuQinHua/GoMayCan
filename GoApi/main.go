package main

import (
	"flag"
	"log"
	"runtime"

	"my_go_api/g"
	"my_go_api/http"
	"my_go_api/store"
)

func prepare() {
	//把可同时执行的最大CPU数目，设置为本机的逻辑CPU数目
	runtime.GOMAXPROCS(runtime.NumCPU())
	//设置日志的输出选项 如何生成用于每条日志的前缀文本
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

//会在main函数被调用前调用
func init() {
	prepare()

	//用名称 c 注册一个配置文件的flag 默认值是cfg.json 返回一个指向flag的指针
	cfg := flag.String("c", "cfg.json", "configuration file")

	//从os.Args[1:](返回除了第一个参数以外的参数)中解析注册的flag  Args保管了命令行参数，第一个是程序名
	flag.Parse()

	handleConfig(*cfg) //解析配置文件

	store.InitMysql() //连接mysql
}

func main() {
	http.Start() //开启http server服务
}

func handleConfig(configFile string) {
	//解析cfg.json的数据 放入结构体GlobalConfig中 然后再赋值给全局变量 Config
	err := g.Parse(configFile)
	if err != nil {
		log.Fatalln(err)
	}
}
