package model

/**
	Author: huang ning
	Date: 2020/01/07
	Func: Redis数据库连接并实现CRUD操作
*/

import (
	"testing"
)

func TestConnectRedis(t *testing.T) {
	Init()
	stringsOps("strlen", "mystr", "")
	listsOps("llen", "mylist", "")
	setsOps("scard", "myset", "")
	zsetsOps("zcard", "myzset",0,"")
	hashedOps("hlen", "myhash", " ", " ")
	//deleteOps( "mynum")
}

