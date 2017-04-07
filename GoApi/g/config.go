package g

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/toolkits/file"
)

type MysqlConfig struct {
	Addr string `json:"addr"`
	Idle int    `json:"idle"`
	Max  int    `json:"max"`
}

type HttpConfig struct {
	Listen string `json:"listen"`
	Secret string `json:"secret"`
}

type GlobalConfig struct {
	Debug bool `json:"debug"`
	// Salt        string              `json:"salt"`
	// Past        int                 `json:"past"`
	Http *HttpConfig `json:"http"`
	// Rpc         *RpcConfig          `json:"rpc"`
	// Log         *LogConfig          `json:"log"`
	Mysql *MysqlConfig `json:"mysql"`
	// Alarm       *AlarmConfig        `json:"alarm"`
	// Falcon      *FalconConfig       `json:"falcon"`
	// InternalDns *InternalDnsConfig  `json:"internalDns"`
	// MonitorMap  map[string][]string `json:"monitorMap"`
}

var (
	Config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

//解析cfg.json的数据 放入结构体GlobalConfig中  然后再赋值给全局变量 Config
func Parse(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		return fmt.Errorf("configuration file %s is nonexistent", cfg)
	}

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		return fmt.Errorf("read configuration file %s fail %s", cfg, err.Error())
	}
	//GlobalConfig 是一个结构体
	var c GlobalConfig
	//解析json编码的数据configContent,并将结果存入c指向的值
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		return fmt.Errorf("parse configuration file %s fail %s", cfg, err.Error())
	}

	configLock.Lock()
	defer configLock.Unlock()
	Config = &c

	log.Println("load configuration file", cfg, "successfully")
	return nil
}
