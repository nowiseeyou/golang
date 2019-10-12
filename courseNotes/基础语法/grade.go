package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(returnStr())
	//grade(101)
	fmt.Println(grade(101))

}

//switch
func grade(score int) string {

	g := ""
	switch {
		case  score < 0 || score >100:
			//panic(fmt.Sprintf("Wrong score %d",score))
			g = "金额错误，超出0-100的范围：" + string(score)
		case  score < 60:
			g = "D"
		case  score < 80:
			g = "C"
		case  score < 90:
			g = "B"
		case  score < 100:
			g = "A"
 	}
	return g
}

//函数返回
func returnStr () string{
	return "ig np"
}
