package main

import (
	"fmt"
	"reflect"
)

/**
	Author: huangning
	Date: 2020/12/30
	Target: 了解slice(引用类型)
	Note: 数组和slice的区别
	1. 前者定长，后者是变长
	2. 前者值类型，把一个数组赋予给另一个数组时是发生值拷贝，
       切片是指针类型，拷贝的是指针, 所以在方法中即使传递切片，其实也是传递的指针。
	3. 添加元素后长度超过容量，那么容量大小自动翻倍
*/
func main() {

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 1.slice的创建
	// make([]T, len, cap)方法创建slice，T是声明的切片保存的数据的类型
	// 其中len表示切片的长度，cap表示capacity
	// 一般使用时省略cap参数，默认和len参数相同
	// capacity和length在切片中是可能不一样的
	// 1) 通过make创建
	info := make([]string, 0, 4)
	// 和[]string{"", "", "", ""}相同

	// 2) 通过数组创建
	sli := arr[0:]
	fmt.Println(info)
	fmt.Printf("arr的类型: %v, sli的类型: %v\n", reflect.TypeOf(arr).Kind(), reflect.TypeOf(sli).Kind())

	// 2.slice增加元素/切片
	sli2 := append(sli, 11, 12, 13)
	// 通过append方法返回的切片是新的切片
	sli3 := append(sli[0:2], sli[3:5]...)
	fmt.Printf("sli: %v 长度: %v 容量: %v\n", sli, len(sli), cap(sli))
	fmt.Printf("新增3个元素的sli2: %v 长度: %v 容量: %v\n", sli2, len(sli2), cap(sli2))
	fmt.Printf("新增2个切片的sli2: %v 长度: %v 容量: %v\n", sli3, len(sli3), cap(sli3))

	// 修改数组/切片
	modifyArr(&arr)
	modifySlice(sli2)
	fmt.Printf("修改后数组的第一个元素为%v\n", arr[0])
	fmt.Printf("修改后切片的第一个元素为%v\n", sli2[0])

	// 切片操作返回的切片与原切片共享存储, 不会复制原来的数据, 而是将新的切片的指针指向原切片的某个元素
	// 拷贝处理长度不同的slice时，只会拷贝较短的部分，将后者值复制给前者
	as := sli2[2:6]
	fmt.Println("copy前as和sli2的值为", as, sli2)
	copy(sli2, as)
	fmt.Println("copy后as的值为", sli2)

}

func modifyArr(a *[10]int)  {
	(*a)[0]=999

}
func modifySlice(a []int)  {
	a[0]=888
}




