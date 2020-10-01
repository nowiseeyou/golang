package main

import "fmt"

// 追加
func main(){
	x := []int{1,2,3}
	x = append(x,4,5,6)
	fmt.Println(x)
	
	x2 := []int{1,2,3}
	y2 := []int{4,5,6}
	
	x2 =  append(x2,y2...) // 如果 没有 ... 它就会由于类型错误而无法编译，因为 y 不是int 类型的
	
	fmt.Println(x2)
}
