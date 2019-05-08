package main

import "github.com/cccCoding/go-practice/day9/example5_chatroom/server/model"

var (
	mgr *model.UserMgr
)

func initUserMgr() {
	mgr = model.NewUserMgr(pool)
}