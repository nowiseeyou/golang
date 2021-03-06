package main

import "fmt"

func main() {
	var numbers []int
	printSlice(numbers)
	
	// 允许追加空切片
	numbers = append(numbers, 0)
	printSlice(numbers)
	
	// 向切片添加一个元素
	numbers = append(numbers, 1)
	printSlice(numbers)
	
	// 同时添加多个元素
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)
	
	// 创建切片number1 是之前切片的两倍容量
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	
	// 拷贝 numbers 的内容到 numbers1
	copy(numbers1, numbers)
	printSlice(numbers1)
	
	// 验证cap计算规则 减掉下限之前的元素
	numbers2 := numbers1[1:6]
	printSlice(numbers2)
}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
