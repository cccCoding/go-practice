package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	Age int
}

func main() {
	user1 := &User{
		UserName:"user1",
		Age:18,
	}
	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("Marshal err:", err)
	}
	fmt.Printf("%s\n", string(data))

	var user2 User
	err = json.Unmarshal(data, &user2)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
	}
	fmt.Println(user2)
}