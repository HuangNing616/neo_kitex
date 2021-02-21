package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)


// 系统运行环境
const (
	Prod  = "prod"
	Dev   = "dev"
	Bench = "bench"
)

// Config 配置信息
type Config struct {
	Env                   string           `json:"env"`
	SiteMeta              SiteMeta         `json:"site"`
	DefaultAbaseTable string           	   `json:"default_abase_table"`
	DefaultAbaseCluster   string           `json:"default_abase_cluster"`
	Accounts              [] *Account      `json:"accounts"`
	AbaseConfigs          [] *AbaseConfig  `json:"abase_configs"`
	BGTable               string           `json:"bg_table"`
	MergeLogKafka         MergeLogKafka    `json:"merge_log_kafka"`
	AccountMap map[string]*Account
}


type SiteMeta struct {
	Host              string `json:"host"`
	RootPath          string `json:"rootPath"`
	SsoHost           string `json:"ssoHost"`
	CookieAgeInSecond int    `json:"cookieAgeInSecond"`
}

type MergeLogKafka struct {
	Cluster string `json:"cluster"`
	Topic   string `json:"topic"`
	Psm     string `json:"psm"`
}

type AbaseConfig struct {
	AbaseTables []string `json:"abase_tables"`
	AbaseCluster string  `json:"abase_cluster"`
}

type Account struct {
	EnName   string `json:"en_name"`
	CnName   string `json:"cn_name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}


var (
	// ConfigInstance 当前环境配置信息
	AccountMap map[string]*Account
)


// Init 初始化配置文件
func Init(file string) error {

	// logs.Info("config file: %s", file)

	// 打开配置文件
	content, err := ioutil.ReadFile("../input/conf_risk_personas_backend_dev.json")
	if err != nil {
		fmt.Println("打开文件报错", err)
	}
	fmt.Printf("%s\n", content)
	var conf Config

	err = json.Unmarshal(content, &conf)
	if err != nil {
		log.Fatalf("init conf instance error:%+v\n", err)
		return err
	}

	// 如果没有配置env，默认是开发环境
	if len(conf.Env) == 0 {
		conf.Env = Dev
	}


	fmt.Println("结构体内容为", conf)
	//// 将account数组中的值取出来
	//parseAccountMap(conf)
	//
	//ConfigInstance = conf
	//SiteMetaInstance = conf.SiteMeta
	return nil

}


func parseAccountMap(conf *Config) {

	AccountMap = make(map[string]*Account)
	// 遍历结构体中的所有account元素, 将其中的密码取出来
	for _, account := range conf.Accounts {
		AccountMap[account.EnName] = account
	}
}

