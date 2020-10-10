package main

import "fmt"

// 切片 切片基于数组构建

// 切片写法 []T, T是切片元素的类型，和数组不同，切片类型并没有固定的长度

// letters := []string{"a","b","c","d"}

// func make([]T,len,cap) []T
// 其中T代表被创建的切片元素的类型。函数 make 接收一个类型、长度和可选容量参数。
// 调用 make 时，内部会分配一个数组，然后返回数组对应的切片
// 当容量参数被忽略，它默认为指定的长度

func main()  {
	letters := []string{"a","b","c","d"}
	fmt.Printf("切片letters : %s \n ", letters)
	
	var s []byte
	s = make([]byte, 5, 5)
	fmt.Printf("切片s : %s \n ", s)
}