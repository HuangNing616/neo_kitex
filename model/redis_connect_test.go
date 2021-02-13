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
	stringsOps("strlen", "mynum", "")
	listsOps("llen", "mylist", "")
	setsOps("scard", "myset", "")
	zsetsOps("zcount", "myzset",0,"", "(1", "inf")
	hashedOps("hlen", "myhash", " ", " ")
	deleteOps( "mynum")

}

