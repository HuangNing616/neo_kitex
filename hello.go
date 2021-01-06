package main

import (
	"github.com/koding/kite"
)

func main() {

	// 创建Kite
	// name = first 并且 version =  1.0.0.
	k := kite.New("math", "1.0.0")
	// Kite配置基本信息
	k.Config.Username = "huangning"
	k.Config.Environment = "production"
	k.Config.Region = "Beijing"
	k.Config.IP = "127.0.0.1"
	k.Config.Id = "2018314003"
	k.Config.Port = 3636

	// 将名为 "square" 的方法添加到我们的 handler method
	k.HandleFunc("square", func(r *kite.Request) (interface{}, error) {
		a := r.Args.One().MustFloat64()
		result := a * a    // calculate the square
		return result, nil // send back the result
	}).DisableAuthentication()

	// 通过 Run() 方法启动了服务器，现在kite就可以接收请求
	// 如果没有分配端口号，操作系统自动为我们选择
	k.Run()
}
