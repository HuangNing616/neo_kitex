package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/**
	Author: huangning
	Date: 2020/12/28
	Funct: 了解string
	Note:
	1. 在go中string类型都是utf-8编码
	2. len统计字节长度, utf.8统计字符长度
*/
func main() {

	//1. 字符串创建
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	name := "Hello Go"
	sex := "男"

	fmt.Println("将字节数组转成字符串:", string(byteSlice))
	fmt.Println("用 len 统计sex的字节数长度:", len(sex))
	fmt.Println("用 utf.8 统计sex的字符数长度:", utf8.RuneCountInString(sex))

	//2. 字符串遍历
	for i:=0;i<len(name);i++ {
		fmt.Printf("字符串 %v, 类型: %T \n", string(name[i]), string(name[i]))
	}

	//3. 判断是否包含某个字符串
	if strings.Contains(name, "Go"){
		fmt.Println("name中包含 World")
	}
}

