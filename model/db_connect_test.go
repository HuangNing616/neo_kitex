package model

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	InitDB()
	Query()
	//Insert(402, "huangxiaoning")
	Delete(301)
	Delete(400)
	Delete(401)
	Delete(402)
}



