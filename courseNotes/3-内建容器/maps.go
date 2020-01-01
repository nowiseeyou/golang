package main

import "fmt"

func main(){

	// 先声明map
	var country map[string]string
	// 再使用make函数创建一个非nil的map，nil map不能赋值
	country = make(map[string]string)
	// 最后给已声明的map赋值

	country["japan"] = "东京"
	country["china"] = "北京"
	country["France"] = "巴黎"

	fmt.Println(country)

	//直接创建

	m2 :=make(map[string]string)
	m2["a"] = "aaa"

	fmt.Println(m2)

	//初始化 并赋值
	m3 := map[string]string{
		"3a":"3333aaaa",
		"3b":"3333bbbb",
	}

	fmt.Println(m3)
	//判断键值是否存在

	v,c := m2["b"]
	fmt.Println(v,c)

	if v,ok := m2["c"];ok{
		fmt.Println(v)
	}else{
		fmt.Println("Key Not Found")
	}

	delete(country,"japan")
	fmt.Println(country)

	for key,value := range country{
		fmt.Println(key,value)
	}
}
