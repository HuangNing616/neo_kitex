package main

import (
	"github.com/koding/kite"
)

/**
	Author: huang ning
	Date: 2021/01/06
	Func: 启动服务器
 */

func main() {

	// 创建Kite
	k := kite.New("huang_server", "1.0.0")

	// 将名为 "square" 的方法添加到我们的 handler method
	k.HandleFunc("square", func(r *kite.Request) (interface{}, error) {
		a := r.Args.One().MustFloat64()
		result := a * a
		return result, nil
	}).DisableAuthentication()

	// 如果没有分配端口号，操作系统自动设置
	k.Config.Port = 3636

	// 通过Run()方法启动服务器，之后kite可以接收请求
	k.Run()
}
