package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	gradeRst := ""
	gradeRst = grade(100)
	fmt.Printf("%s \n",gradeRst)
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
