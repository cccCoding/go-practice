package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)


func main() {
	//链接redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Error redis Dial", err)
		return
	}
	defer conn.Close()

	//Set接口
	_, err = conn.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println("error Set", err)
		return
	}

	//Get
	r, err := redis.Int(conn.Do("Get", "abc"))
	if err != nil {
		fmt.Println("Get abc failed", err)
		return
	}
	fmt.Println(r)
}
