package main

import "fmt"

/**
	Author: huangning
	Date: 2020/12/30
	Target: 了解recover
	Note:
	1. go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理
	2. 利用recover处理panic指令，defer必须在panic之前声明，否则当panic时，recover无法捕获到panic
*/
func main() {
	fmt.Println("c")
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
		fmt.Println("e")
	}()
	f() //开始调用f
	fmt.Println("f") //这里开始下面代码不会再执行
}

func f() {

	fmt.Println("a")
	panic("异常信息")
	fmt.Println("b") //这里开始代码不会再执行

}
