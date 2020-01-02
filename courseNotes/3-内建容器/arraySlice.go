package main

import "fmt"

func main(){
	arr1 := [...]int{1,3,5,7,9,11}
	slice1 := arr1[:]
	slice2 := arr1[:3] //[1,3,5]
	slice3 := slice2[1:4] //[3,5,7]

	fmt.Println("arr1")
	fmt.Println(arr1)

	fmt.Println("slice1")
	setArr0(slice1)

	fmt.Println("arr1")
	fmt.Println(arr1)

	fmt.Println("slice3")
	fmt.Println(slice3)
	fmt.Printf("slice3 len:%d , cap:%d ,self:%v \n",len(slice3),cap(slice3),slice3)


}

func setArr0(arr []int){
	arr[0] = 100
	fmt.Println(arr)
}

