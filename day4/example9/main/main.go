package main

import "fmt"

func testMap() {
	//var a map[string]string			//声明，后面必须make分配内存进行初始化
	//a = make(map[string]string)		//分配内存空间初始化

	//a := make(map[string]string)		//声明并初始化

	var a = map[string]string{			//声明并初始化
		"key": "value",
	}

	a["abc"] = "我爱你中国"
	a["abc"] = "我爱你母亲"		//覆盖相同key的value
	a["cba"] = "嘻嘻"
	fmt.Println(a)
}

func testMap2() {
	a := make(map[string]map[string]string)
	a["key"] = make(map[string]string, 5)	//初始化先
	a["key"]["key1"] = "value1"
	a["key"]["key2"] = "value2"
	a["key"]["key3"] = "value3"
	fmt.Println(a)

	delete(a["key"], "key3")			//删除
	fmt.Println(a)

	//a["keykey"] = make(map[string]string, 5)	//初始化先
	//a["keykey"]["key1"] = "value1"
	//a["keykey"]["key2"] = "value2"
	//a["keykey"]["key3"] = "value3"
	_ ,ok := a["keykey"]
	if !ok {
		a["keykey"] = make(map[string]string, 5)	//如果value为map类型且没有初始化，要先初始化
	}
	a["keykey"]["key2"] = "value22"		//替换	//记住这种写法

	fmt.Println(a)
}

func main() {
	testMap()
	testMap2()
}
