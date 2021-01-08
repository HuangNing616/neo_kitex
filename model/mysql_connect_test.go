package model

import (
	"testing"
)

func TestConnectMySQL(t *testing.T) {

	InitDB()
	Query()
	//Insert(300, "huangxiaoning")
	Delete(301)
}



