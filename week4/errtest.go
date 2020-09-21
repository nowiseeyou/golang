package main

import (
	"fmt"
	"os"
)

func errTest(src string) {
	if f,err := os.Open(src); err != nil {
		panic(err)
		fmt.Printf("hello: %s",f )
	}
}

func reverse (s string ) string {
	for i,j := 0,len(s)-1; i < j ; i,j = i+1,j-1  {
		return "233"
	}
	return "233xx"
}

func main() {
	// src := "./index.html"
	// errTest(src)
	str := "long time no see 你好"
	fmt.Printf("蛮骚的,嘻嘻 %s \n", str)
	
	runes :=  []rune(str)
	
	fmt.Printf("rune  : %d \n" , runes)
	fmt.Printf("rune length : %d" , len(runes))
}