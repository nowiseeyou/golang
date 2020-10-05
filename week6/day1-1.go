package main

import (
	"fmt"
	"os"
)

// 空白标示符 _
// 若某次赋值需要匹配多个左值，但其中某个变量不会被程序使用， 那么用空白标示符来代替该变量可避免创建无用的变量，
// 并能清楚地表明该值将被丢弃。


func main(){
	path := "/jayzhou-qlx.txt"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("%s does not exist \n", path)
		// fmt.Printf("Content : %s  \n", content)
	}
}
