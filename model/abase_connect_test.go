package model

/**
	Author: huang ning
	Date: 2020/01/07
	Func: Redis数据库连接并实现CRUD操作
*/

import (
	"fmt"
	"testing"
)

func TestConnectAbase(t *testing.T) {


	fmt.Println("连接Abase... ...")
	//options := goredis.NewOption()
	//// 目前正在推取消ss_conf的使用，后续将不维护ss_conf，所以不要指定ss_conf路径
	//options.IdleTimeout = time.Minute * 30
	//client, err := goredis.NewClientWithOption("ttt", options)
	//if err != nil {
	//	log.Println("new client err: ", err)
	//}
	//
	//// example
	////表名为'zs_test', key为'hldd'
	//ret := client.Get("[zs_test]hldd")
	//log.Println("ret is ", ret)

}

