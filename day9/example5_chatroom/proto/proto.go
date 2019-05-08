package proto

import "github.com/cccCoding/go-practice/day9/example5_chatroom/model"

type Message struct {
	Cmd string `json:"cmd"`
	Data string `json:"data"`
}

type LoginCmd struct {
	Id int `json:"user_id"`
	Passwd string `json:"passwd"`
}

type RegisterCmd struct {
	User model.User `json:"user"`
}

type LoginCmdRes struct {
	Code int `json:"code"`
	Error string `json:"error"`
}


