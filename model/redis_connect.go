package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"strings"
)

var pool *redis.Pool
const RedisHost = "127.0.0.1:6379"

// 初始化redis连接池
func Init()  {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", RedisHost)
		},
	}
}

func stringsOps(command string, key string, value string, args...string)  {

	command = strings.ToUpper(command)
	stringConn := pool.Get()
	if command == "SET" {
		_, err := stringConn.Do("SET", key, value)
		if err != nil {
			fmt.Println("SET err=", err)
		}
		fmt.Println("set operation success !!!")
	}else if command == "GET" {
		val, err := redis.String(stringConn.Do("GET", key))
		if err != nil {
			fmt.Println("GET err=", err)
		}
		fmt.Printf("key = %v , value = %v \n", key, val)
	}else if command == "STRLEN" {
		length, err := stringConn.Do("STRLEN", key)
		if err != nil {
			fmt.Println("STRLEN err=", err)
		}
		fmt.Printf("The length of <%v> is %v \n", key, length)
	} else{
		log.Fatal("current command has not been supported")
	}
}

func listsOps(command string, key string, value string, args...string)  {

	command = strings.ToUpper(command)
	listConn := pool.Get()
	if command == "LPUSH"{
		_, err := listConn.Do("LPUSH", key, value)
		if err != nil {
			fmt.Println("LPUSH err=", err)
		}
		fmt.Println("LPUSH operation success !!!")
	}else if command == "RPUSH"{
		_, err := listConn.Do("RPUSH", key, value)
		if err != nil {
			fmt.Println("RPUSH err=", err)
		}
		fmt.Println("rpush operation success !!!")
	}else if command == "LPOP"{
		elem, err := redis.String(listConn.Do("LPOP", key))
		if err != nil {
			fmt.Println("LPOP err=", err)
		}

		fmt.Printf("first elements of the list = %v\n", elem)
		fmt.Println("lpop operation success !!!")
	}else if command == "RPOP"{
		elem, err := redis.String(listConn.Do("RPOP", key))
		if err != nil {
			fmt.Println("RPOP err=", err)
		}
		fmt.Printf("last elements of the list = %v\n", elem)
		fmt.Println("rpop operation success !!!")
	}else if command == "LRANGE"{
		values, err := redis.Values(listConn.Do("LRANGE", key, args[0], args[1]))
		if err != nil {
			fmt.Println("LRANGE err=", err)
		}
		for index, v := range values {
			fmt.Printf("第%v个元素=%s \n", index, v)
		}
		fmt.Println("LRANGE operation success !!!")
	}else if command == "LLEN"{
		res, err := listConn.Do("LLEN", key)
		if err != nil {
			fmt.Println("LLEN err=", err)
		}
		fmt.Printf("The length of <%v> is %v\n", key, res)
	}else{
		log.Fatal("current command has not been supported")
	}
}

func setsOps(command string, key string, value string)  {
	command = strings.ToUpper(command)
	stringConn := pool.Get()
	if command == "SADD" {
		_, err := stringConn.Do("SADD", key, value)
		if err != nil {
			fmt.Println("SADD err=", err)
		}
		fmt.Println("SADD operation success !!!")
	}else if command == "SCARD"{
		count, err := stringConn.Do("SCARD", key)
		if err != nil {
			fmt.Println("SCARD err=", err)
		}
		fmt.Printf("The number of <%v> is %v\n", key, count)
	}else{
		log.Fatal("current command has not been supported")
	}
}

func zsetsOps(command string, key string, index int, value string)  {

	command = strings.ToUpper(command)
	stringConn := pool.Get()
	if command == "ZADD" {
		_, err := stringConn.Do("ZADD", key, index, value)
		if err != nil {
			fmt.Println("ZADD err=", err)
		}
		fmt.Println("ZADD operation success !!!")
	}else if command == "ZCARD"{
		count, err := stringConn.Do("ZCARD", key)
		if err != nil {
			fmt.Println("ZCARD err=", err)
		}
		fmt.Printf("The number of <%v> is %v\n", key, count)
	}else{
		log.Fatal("current command has not been supported")
	}
}

func hashedOps(command string, key string, filed string, value string)  {

	command = strings.ToUpper(command)
	stringConn := pool.Get()
	if command == "HSET" {
		_, err := stringConn.Do("HSET", key, filed, value)
		if err != nil {
			fmt.Println("HSET err=", err)
		}
		fmt.Println("HSET operation success !!!")
	}else if command == "HGET"{
		val, err := redis.String(stringConn.Do("HGET", key, filed))
		if err != nil {
			fmt.Println("HGET err=", err)
		}
		fmt.Printf("key=<%v>, field = %v, value=%v\n", key, filed, val)
	}else if command == "HLEN"{
		length, err := stringConn.Do("HLEN", key)
		if err != nil {
			fmt.Println("HLEN err=", err)
		}
		fmt.Printf("The length of <%v> is %v\n", key, length)
	} else{
		log.Fatal("current command has not been supported")
	}
}


func deleteOps(key string)  {

	delConn := pool.Get()
	count, err := delConn.Do("DEL", key)
	if err != nil {
		fmt.Println("DEL err=", err)
	}
	fmt.Printf("DEL elements count = %v\n", count)
}