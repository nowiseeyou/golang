package adder

import "fmt"

func adder(int) func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main(){
	a := adder(1)
	for i := 0; i <10 ; i++ {
		fmt.Println("0 + 1 + .... + %d = %d \n",i,a(i))
	}
}

