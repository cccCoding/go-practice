package main

import (
	"fmt"
	"net"
)

// 了解 Nginx redis
// telnet localhost 50000	????

// 服务端处理流程 1，监听端口 2，接收客户端的链接 3，创建goroutine，处理该链接
// 客户端处理流程 1，建立与服务端的链接 2，进行数据收发 3，关闭链接

func main() {
	fmt.Println("start server...")
	listen, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}