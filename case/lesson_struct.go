package main

import (
	"fmt"
	"reflect"
)

/**
	Author: huangning
	Date: 2020/12/28
	Target: 了解struct(值类型)
	Note:
	在一个结构体中对于每一种数据类型只能有一个匿名字段
*/

type Person struct {
	Name string `Configuration:"-"`
	Age  int    `Configuration:"age,min=17,max=60"`
	Sex  string `Configuration:"sex,required"`
	Email string `Configuration:"Email,required @xxx sign"`
}

// 学生结构体
type Student struct {
	grade, height float32
	Person  // 内嵌结构体
	bool  // 内嵌(匿名)字段
}

func main() {

	// 1. 初始化结构体类型的变量
	// 第1种结构体字面量
	var p1 = Person{"zhang_san", 38, "man", "huang@qq.com"}

	// 第2种结构体字面量
	p2 := Student{
		grade: 78.2,
		height: 178.2,
		Person: p1,
		bool: true,
	}

	// 2.结构体没有手动初始化，系统也会默认初始化
	var nick Person
	fmt.Println("默认的姓名:", nick.Name)

	// 3.结构体元素访问, 通过选择器selector, 即.操作符号
	fmt.Printf("结构体p1的姓名:%v \n", p1.Name)
	fmt.Printf("结构体p2的成绩:%v \n", p2.grade)

	// 4.结构体元素修改
	fmt.Printf("p2修改前的身高:%v \n", p2.grade)
	p2.grade = 99.9
	fmt.Printf("p2修改后的身高:%v \n", p2.grade)

	// 5.通过反射，获取结构体变量的动态类型
	const tagName = "Configuration"
	strut := reflect.TypeOf(p1)
	fmt.Printf("结构体p1的名称: %v 类型: %v \n", strut.Name(), strut.Kind())
	for i:=0; i<strut.NumField(); i++ {
		field := strut.Field(i)
		fmt.Printf("第%d个字段名称:【%v】, 字段类型名称:【%v】, 标签tag:【%v】\n", i+1, field.Name, field.Type.Name(), field.Tag.Get(tagName))
	}

	// 6. 用new方法初始化变量
	// 把内存初始化为零值并返回其指针
	P1 := new(Person)
	var P2 Person
	fmt.Printf("new出来的变量类型是: %T \n", P1)
	fmt.Printf("声明出来的变量类型为: %T \n", P2)

}

