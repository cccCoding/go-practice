package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cccCoding/go-practice/day9/example5_chatroom/proto"
	"net"
)

type Client struct {
	conn net.Conn
	buf [8192]byte
}

func (c *Client) Process() (err error) {
	for {
		var msg proto.Message
		msg, err = c.readPackage()
		if err != nil {
			return err
		}
		err = c.processMsg(msg)
		if err != nil {
			return
		}
	}
}

func (c *Client) readPackage() (msg proto.Message, err error) {
	n, err := c.conn.Read(c.buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}

	var packLen = binary.BigEndian.Uint32(c.buf[0:4])

	n, err = c.conn.Read(c.buf[0:packLen])
	if err != nil || n != int(packLen) {
		err = errors.New("read data failed")
		return
	}

	err = json.Unmarshal(c.buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed", err)
	}
	return
}

func (c *Client) writePackage(data []byte) (err error) {
	packLen := uint32(len(data))
	binary.BigEndian.PutUint32(c.buf[0:4], packLen)
	n, err := c.conn.Write(c.buf[0:4])
	if err != nil {
		fmt.Println("write data len failed")
		return
	}

	n, err = c.conn.Write(data)
	if err != nil {
		fmt.Println("write data failed")
		return
	}

	if n != int(packLen) {
		fmt.Println("write data not finished")
		err = errors.New("write data not finished")
		return
	}

	return
}

func (c *Client) processMsg(msg proto.Message) (err error) {
	switch msg.Cmd {
	case proto.UserLogin:
		err = c.login(msg)
	case proto.UserRegister:
		err = c.register(msg)
	default:
		err = errors.New("unsupported message")
	}
	return
}

func (c *Client) login(msg proto.Message) (err error) {
	defer func() {
		c.loginResp(err)
	}()
	fmt.Printf("recv user login request, data:%v\n", msg)
	var cmd proto.LoginCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}
	_, err = mgr.Login(cmd.Id, cmd.Passwd)
	if err != nil {
		return
	}
	return
}

func (c *Client) loginResp(err error) {
	var respMsg proto.Message
	respMsg.Cmd = proto.UserLoginRes

	var loginRes proto.LoginCmdRes
	loginRes.Code = 200
	if err != nil {
		loginRes.Code = 500
		loginRes.Error = fmt.Sprintf("%v", err)
	}

	data, err := json.Marshal(loginRes)
	if err != nil {
		fmt.Println("marshal loginRes failed", err)
		return
	}
	respMsg.Data = string(data)

	data, err = json.Marshal(respMsg)
	if err != nil {
		fmt.Println("marshal respMsg failed", err)
		return
	}
	err = c.writePackage(data)
	if err != nil {
		fmt.Println("write pkg failed", err)
		return
	}
}

func (c *Client) register(msg proto.Message) (err error) {
	var cmd proto.RegisterCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}
	err = mgr.Register(&cmd.User)
	if err != nil {
		return
	}
	return
}
