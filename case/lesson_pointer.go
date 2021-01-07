package main

import "fmt"

/**
	Author: huangning
	Date: 2020/12/29
	Map: 了解指针
	Note:
	1. 常量无法获取指针
	2. * 放在类型旁边，表示指向这个类型的指针, 放在变量旁边，表示变量指向的值
 */

func changePointerValue(a *int)  {
	*a = 55
}

func main()  {

	var sex string = "man"
	age := 28

	// 1.通过 & 获取地址
	var sexPt *string = &sex
	agePt := &age
	agePP1 := &agePt
	var agePP2 **int = &agePt

	fmt.Printf("变量 sex 的值: %v, 地址: %v\n", sex, sexPt)
	fmt.Printf("变量 age 的值: %v, 地址: %v\n", age, agePt)
	fmt.Printf("sexPt 是指向 string 类型变量的指针, sexPt的类型: %T, 指向的值为%v\n", sexPt, *sexPt)
	fmt.Printf("==================\n")
	fmt.Printf("agePP1 和 agePP2 都是指向变量 age 指针的指针\n")
	fmt.Printf("agePP1的类型: %T, 指向指针所指向的值为%v\n", agePP1, *(*agePP1))
	fmt.Printf("agePP2的类型: %T, 指向指针所指向的值为%v\n", agePP2, *(*agePP2))

	// 2.指针的零值(没有做初始化时系统默认的值)
	var zero *int
	fmt.Printf("指针的零值:%v 类型:%T \n", zero, zero)

	// 3. 修改指针指向的元素
	fmt.Printf("修改前age的值为%v\n", age)
	changePointerValue(agePt)
	fmt.Printf("修改后age的值为%v\n", age)
}
