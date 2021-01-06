package client

import (
	"fmt"
	"github.com/koding/kite"
	"testing"
)

// TestGetAllFeatureConfig
func TestHuang(t *testing.T) {

	k := kite.New("huang_client", "1.0.0")

	// Connect to our math kite
	mathWorker := k.NewClient("http://localhost:3636/kite")
	mathWorker.Dial()

	response, _ := mathWorker.Tell("square", 4) // call "square" method with argument 4

	t.Log("日志信息开始记录... ...")

	fmt.Println("接收到的结果为:", response)
	fmt.Println("将结果转成float:", response.MustFloat64())

	t.Log("TestHuang done")
}


