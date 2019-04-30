package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("len of args:%d\n", len(os.Args))
	//以空格隔开各个参数
	for i, v := range os.Args {
		fmt.Printf("args[%d]:%s\n", i, v)	//args[0]为程序本身
	}

	//使用封装好的flag包来解析命令行参数
	var confPath string
	var logLevel int
	flag.StringVar(&confPath, "c", "", "plz input config path")
	flag.IntVar(&logLevel, "d", 0, "plz input log level")

	flag.Parse()

	fmt.Println("path:", confPath)	//有了路径可以读取配置文件进行初始化
	fmt.Println("logLevel:", logLevel)

	//./main -c testpath -d 100
}