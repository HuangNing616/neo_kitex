package config

import (
	"fmt"
	"testing"
)

func TestConfigInit(t *testing.T) {

	// 打开配置文件
	err := Init("../input/conf_risk_personas_backend_dev.json")
	if err != nil {
		fmt.Println("打开文件报错", err)
	}
}
