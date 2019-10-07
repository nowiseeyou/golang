package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	grade(101)
	//fmt.Printf(grade(101))

}

//switch
func grade(score int) string {

	g := ""
	switch  {
	case  score < 0 || score >100:
		panic(fmt.Sprintf("Wrong score: %d",score))
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
