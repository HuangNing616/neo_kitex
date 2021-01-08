package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestConfigInit(t *testing.T) {

	// 打开配置文件
	content, err := ioutil.ReadFile("../input/conf_risk_personas_backend_dev.json")
	if err != nil {
		fmt.Println("打开文件报错", err)
	}

	var conf Config
	err = json.Unmarshal(content, &conf)
	if err != nil {
		return
	}
}

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
