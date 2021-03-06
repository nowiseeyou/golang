package main

import "fmt"

func main(){
	fmt.Println("Creating slice")
	var s []int //zero value for slice is nil
	for i := 0; i<100 ;i++ {
		s = append(s,2*i +1)
	}
	fmt.Println(s)

	s1 := []int{2,4,6,8}
	printSlice(s1)

	s2 := make([]int,16)
	s3 := make([]int,16,32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2,s1)
	printSlice(s2)

	fmt.Println("Deleting elements form slice")
	s2 = append(s2[:3],s2[4:]...)	//删除key 为3的元素
	printSlice(s2)

	fmt.Println("Poping from slice")
	front := s2[0]
	s2 = s2[1:]				//删除头部元素
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Poping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]		//删除尾部元素
	fmt.Println(tail)
	printSlice(s2)
}

func printSlice(s []int){
	fmt.Printf("%v，len =%d，cap =%d\n",s,len(s),cap(s))
}
