package main

import (
	"encoding/json"
	"fmt"
)

/**
	Author: huang ning
	Date: 2021/01/06
	Func:
 */


type StuRead struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	HIgh  interface{}
	sex   interface{}
	Class interface{} `json:"class"`
	Test  interface{}
}

type Class struct {
	Name  string
	Grade int
}

func main() {
	//json字符中的"引号，需用\进行转义，否则编译出错
	//json字符串沿用上面的结果，但对key进行了大小的修改，并添加了sex数据
	data:="{\"name\":\"张三\",\"Age\":18,\"high\":true,\"sex\":\"男\",\"CLASS\":{\"naME\":\"1班\",\"GradE\":3}}"
	str:=[]byte(data)

	//1.Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构。
	//第二个参数必须是指针，否则无法接收解析的数据，如stu仍为空对象StuRead{}
	//2.可以直接stu:=new(StuRead),此时的stu自身就是指针
	stu:=StuRead{}
	err:=json.Unmarshal(str,&stu)

	//解析失败会报错，如json字符串格式不对，缺"号，缺}等。
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println(stu)
}

//type Stu struct {
//	Name  string `json:"hnname"`
//	Age   int
//	HIgh  bool
//	sex   string
//	Class *Class `json:"hnclass"`
//}
//
//type Class struct {
//	Name  string
//	Grade int
//}
//
//func main() {
//
//	//实例化一个数据结构，用于生成json字符串
//	stu := Stu{
//		Name: "张三",
//		Age:  18,
//		HIgh: true,
//		sex:  "男",
//	}
//
//	//指针变量
//	cla := new(Class)
//	cla.Name = "1班"
//	cla.Grade = 3
//	stu.Class=cla
//
//	//Marshal失败时err!=nil
//	jsonStu, err := json.Marshal(stu)
//	if err != nil {
//		fmt.Println("生成json字符串错误")
//	}
//
//	//jsonStu是[]byte类型，转化成string类型便于查看
//	fmt.Println(string(jsonStu))
//}

//func main() {
//
//	k := kite.New("nep_server", "1.0.0")
//	k.HandleFunc("square", func(r *kite.Request) (interface{}, error) {
//		a := r.Args.One().MustFloat64()
//		result := a * a
//		return result, nil
//	}).DisableAuthentication()
//	k.Config.Port = 3636
//	k.Run()
//
//}
