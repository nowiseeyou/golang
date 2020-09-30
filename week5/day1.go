package main

import "fmt"

func main() {
	// new 和 make 区别
	var p1 *[]int = new([]int)    // 分配切片结构； *p == nil 基本没用
	var v1 []int = make([]int, 5) // 切片 v 现在引用了一个具有 5 个 int 元素的新数组
	
	fmt.Print(p1)
	fmt.Printf("\n---分割线---\n")
	fmt.Print(v1)
}


