package main

import "fmt"

// 切片 半闭半开区间(]
func Append(slice, data []byte) []byte {
	l := len(slice)
	if l + len(data) > cap(slice) { // 重新分配
		// 为了后边的增长，需要分配两份
		newSlice := make([]byte, (l + len(data)) * 2)
		
		// copy 函数是预声明的，且可用于任何切片类型
		copy(newSlice, slice)
		
		slice = newSlice
	
	}
	
	slice = slice[0:1 + len(data)]
	
	for i,c := range data {
		slice[1+i] = c
	}
	
	return slice
}

func main(){
	type Transform [3][3]float64    // 一个 3x3 的数组，其实是包含多个数组的一个数组
	type linesOfText [][]byte       // 包含多个字节切片的一个切片

	text := linesOfText{
		[] byte("Now is the time"),
		[] byte("for all good gophers"),
		[] byte("to bring some fun to the party"),
	}
	
	fmt.Println(text)
	
	// 1 独立分配每一个切片
	// 分配顶层切片
	picture1 := make([][]uint8, Ysize) //  每 y 个单元一行
	
	// 遍历行，为每一行 都分配切片
	for i := range picture1 {
		picture1[i] = make([]uint8, Xsize)
	}
	
	
	// 2 只分配一个数组
	// 分配顶层切片 和前边一样
	picture := make([][]uint8, Ysize) //  每 y 个单元一行
	// 分配一个大的切片来保存所有像素
	pixels := make([]uint8, Xsize * Ysize) // 拥有类型 []uint8, 尽管图片是 [][]uint8
	
	// 遍历行， 从剩余像素切片的前面切出每行来
	
	for i := range picture{
		picture[i], pixels = pixels[:Xsize], pixels[Xsize:]
	}
	
	fmt.Println(picture)
	
 }
