package main

import (
	"fmt"
)

/**
	Author: 黄宁
	Date: 2021/02/15
	Func: 值类型
	Note: 包括字符串、整型、浮点型、布尔型等
 */

func main() {

	// 字符串通过+连接
    fmt.Println("go" + "lang")
	
	// 整数运算
    fmt.Println("1+1 =", 1+1)

	// 浮点数运算
    fmt.Println("7.0/3.0 =", 7.0/3.0)

	// 布尔运算
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}