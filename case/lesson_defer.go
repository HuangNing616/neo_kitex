package main

import "fmt"


/**
	Author: huangning
	Date: 2020/12/30
	Target: 了解defer
	Note:
	1. 关键字defer允许我们推迟到函数返回之前
*/

//
func main() {
	doDBOperations()
}

func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return
	// terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}
