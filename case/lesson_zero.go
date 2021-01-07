package main

import "fmt"

/**
	Author: huangning
	Date: 2020/12/30
	Target: 了解零值(变量没有做初始化时系统默认设置的值)
*/

func main() {

	var aa bool
	var bb string
	var cc *int
	var dd []int
	var ee map[string] int
	var ff chan int
	var gg func(string) int
	var hh error
	fmt.Printf("bool型零值: %v\n", aa)
	fmt.Printf("string型零值: %v\n", bb)
	fmt.Printf("数组型零值: %v\n", dd)
	fmt.Printf("map型零值: %v\n", ee)
	fmt.Printf("指针型零值: %v\n", cc)
	fmt.Printf("chan型零值: %v\n", ff)
	fmt.Printf("func型零值: %v\n", gg)
	fmt.Printf("err型零值: %v\n", hh)
	fmt.Println("其他类型的零值都为0")

}
