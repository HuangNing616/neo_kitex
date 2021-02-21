package main

import (
	"fmt"
)

/**
	Author: 黄宁
	Date: 2021/02/15
	Func: 变量
	Note: <变量>需要显式声明，并且在函数调用等情况下，编译器会检查其类型的正确性
 */

func main() {

    var a = "initial"
    fmt.Println(a)

    // 可以一次性声明多个变量
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go 会自动推断已经有初始值的变量的类型
    var d = true
    fmt.Println(d)

    // 声明后却没有给出对应的初始值时，变量将会初始化为<零值? 例如int的零值是0
    var e int
    fmt.Println(e)

    // := 语法是声明并初始化变量的简写, 等价于 var f string = "short" 可以简写为右边这样
    f := "short"
    fmt.Println(f)
}