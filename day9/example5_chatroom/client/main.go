package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cccCoding/go-practice/day9/example5_chatroom/proto"
	"net"
)

func login(conn net.Conn) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserLogin

	var loginCmd proto.LoginCmd
	loginCmd.Id = 1
	loginCmd.Passwd = "123456"

	data, err := json.Marshal(loginCmd)
	if err != nil {
		fmt.Println("Marshal loginCmd failed", err)
		return
	}

	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("Marshal msg failed", err)
		return
	}

	//先写长度
	var buf [4]byte
	packLen := uint32(len(data))
	binary.BigEndian.PutUint32(buf[:], packLen)

	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write data len failed")
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("write data failed", err)
		return
	}

	msg, err = readPackage(conn)
	if err != nil {
		fmt.Println("readPackage failed", err)
		return
	}
	fmt.Println(msg)
	return
}

func readPackage(conn net.Conn) (msg proto.Message, err error) {
	var buf [8192]byte
	n, err := conn.Read(buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}

	var packLen = binary.BigEndian.Uint32(buf[0:4])

	n, err = conn.Read(buf[0:packLen])
	if err != nil || n != int(packLen) {
		err = errors.New("read data failed")
		return
	}

	err = json.Unmarshal(buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed", err)
	}
	return
}

func main() {
	conn, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("dialing failed", err)
		return
	}

	err = login(conn)
	if err != nil {
		fmt.Println("login failed", err)
		return
	}
}
