package main

import "fmt"

func main() {
	var a int = 20 // 声明实际变量
	var ip *int    // 声明指针变量
	
	ip = &a // 指正变量存储地址
	
	fmt.Printf("a 变量地址是：%x\n", &a)
	
	fmt.Printf("ip 变量储存的指针地址：%x\n", ip)
	
	// ip变量的指针访问值
	fmt.Printf("ip 变量的值：%d\n", *ip)
}
