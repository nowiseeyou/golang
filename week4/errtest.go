package main

import (
	"fmt"
	"os"
)

// if err
func errTest(src string) {
	if _, err := os.Open(src); err == nil {
		panic(err)
	}
}

//  字符串反序输出
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
		fmt.Printf("i: %d, j: %d \n", runes[i], runes[j])
	}
	return string(runes)
}

// switch grade
func grade(score float32) string {
	switch {
	case score > 50, score >= 100:
		return "D"
	case score > 70:
		return "C"
	case score > 80:
		return "B"
	case score > 90, score < 80:
		return "A"
	}
	return "N"
}

/**
Compare 按字典顺序比较两个字节切片并返回一个整数
若 a == b 则结果为 ： 0 ； 若 a < b  则结果 ： -1 ； 若 a > b  则结果 ： 1
*/
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func main() {
	grade1 := grade(88)
	grade2 := grade(101)
	src := "./index.html"
	errTest(src)
	fmt.Printf("grade: %s \n", grade1)
	fmt.Printf("grade2: %s \n", grade2)
	
	str := "long time no see 你好"
	fmt.Printf("蛮骚的,嘻嘻 %s \n", str)
	
	reverses := reverse(str)
	fmt.Printf("reverses : %s \n", reverses)
	
	runes := []rune(str)
	fmt.Printf("rune  : %d \n", runes)
	fmt.Printf("rune length : %d \n", len(runes))
}
