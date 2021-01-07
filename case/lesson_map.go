package main

import (
	"fmt"
)

/**
	Author: huangning
	Date: 2020/12/28
	Target: 了解map(引用类型)
	Note:
	1. map是引用类型
	2. map的key值必须可hash
	3. nil和empty的map不同，前者是空指针，而后者指向的内存为0
	4. 用slice保存map中的key，需要保证key和slice中的数据类型一致
 */
func main()  {

	//1. map创建
	studentGrade := map[string]float32{"Huang":100, "Tom":98.2, "John":71.6}
	teacherSalary := make(map[string]int)
	//使用函数func作为map的value
	personSex:=map[string]func()string{
		"zhang_san": func() string {
			return "man"
		},
		"li_si": func() string {
			return "woman"
		},
		"zhao_wu": func() string {
			return "man"
		},
	}

	//2. map插入
	studentGrade["Bob"] = 84.7
	teacherSalary["MrLiu"] = 5000
	teacherSalary["MissWang"] = 20000
	teacherSalary["MrsZhang"] = 12000

	//3. map删除
	delete(studentGrade, "Huang")

	//4. map修改
	teacherSalary["MissWang"] = 18000

	//5. map查询
	fmt.Println("学生的成绩信息:", studentGrade)
	fmt.Println("其中Bob的成绩:", studentGrade["Bob"])
	fmt.Println("zhang_san的性别:", personSex["zhang_san"]())

	//6. map遍历
	for name, salary:=range teacherSalary{
		fmt.Printf("当前老师的姓名: %v, 薪资为: %v \n", name, salary)
	}
	for _, fun:=range personSex{
		fmt.Printf("当前用户的性别: %v \n", fun())
	}

	//7. 判断map是否相等
	//方法1
	var nilDemo map[string]int
	if nilDemo == nil{
		fmt.Println("当前map类型为nil")
	}

	//方法2
	studentGrade2 := studentGrade
	if len(studentGrade) !=len(studentGrade2) {
		fmt.Println("两个map长度不同，二者必然不相同")
	}else{
		for name, grade:= range studentGrade2{
			if studentGrade2[name]!=studentGrade[name] {
				break
			}
			fmt.Printf("当前姓名为%v的成绩相同, 前者的成绩%v, 后者的成绩为%v, \n", name, grade, studentGrade[name])
		}
	}

	//8. 判断map是是否存在指定的key
	//方法1, 根据map单返回值是否系统默认值0, 风险在于有些key对应的value就是0
	if grade1 := studentGrade["Zoe"]; grade1==0 {
		fmt.Println("Zoe这个人不再学生信息map里面")
	}
	//方法2, 根据map双返回值的第二个值
	if grade2, exists := studentGrade["Tim"]; exists==false {
		fmt.Printf("Tim这个人不再学生信息map里面, 对应的系统默认值为%v, \n", grade2)
	}

	//9. 保存map的key
	keyList := make([]string, 3, len(studentGrade))
	fmt.Println("初始化后key_list:", keyList, "长度:", len(keyList), "容量:", cap(keyList))
	for name, _ := range teacherSalary{
		keyList = append(keyList, name)
	}
	keyList = append(keyList, "final")
	fmt.Println("添加map的key后key_list:", keyList, "长度:", len(keyList), "容量:", cap(keyList))
	fmt.Println("从上述结果可以看出，slice元素超出容量之后，新的容量会变成原来的2倍")





}
