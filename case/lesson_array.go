package main

import (
	"fmt"
	"reflect"
)

/**
	Author: huangning
	Date: 2020/12/30
	Target: 了解数组
	Note:
	1. 定义：是一组同类型元素的序列
	2. array存储元素的类型+长度，alice只存元素类型
	3. array的重要特性是长度不可变, slice长度可变
	4. 数组是值类型，把一个数组赋予给另一个数组时是发生值拷贝，而切片是指针类型，拷贝的是指针
	5. 在括号中定义长度值，是编译器区别变量一依据，带长度值便是数组，不带便是切片
*/
func main()  {

	// 1. 数组初始化
	arr1:=[3]string{}
	arr2:=[6]string{}
	var arr3 [4]int
	// 定义直接确定长度，而是由实际元素个数确定长度
	arr4 := [...]string{"Peter", "John", "Mike"}

	// 不同长度也代表不同的数据类型
	fmt.Printf("type of arr1 is %s, the kind is %s\n",reflect.TypeOf(arr1),reflect.TypeOf(arr1).Kind())
	fmt.Printf("type of arr2 is %s, the kind is %s\n",reflect.TypeOf(arr2),reflect.TypeOf(arr2).Kind())
	fmt.Printf("type of arr3 is %s, the kind is %s\n",reflect.TypeOf(arr3),reflect.TypeOf(arr3).Kind())
	fmt.Printf("type of arr4 is %s, the kind is %s\n",reflect.TypeOf(arr4),reflect.TypeOf(arr4).Kind())






}
