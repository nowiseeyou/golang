package main

import "fmt"

func main() {
	var n [10]int // n是一个长度为10的数组
	var i, j int
	
	// 为数组 n 初始化元素
	for i = 0; i < 10; i++ {
		n[i] = 100 + i // 设置元素为   100 + i
	}
	
	// 输出每个数组元素的值
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}
