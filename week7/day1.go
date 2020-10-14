package main

import "fmt"

// 冒泡算法

func bubbleSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len; i++ { // 1-4 2-3 3-2 2-1 1
		for j := 0; j < len-1-i; j++ { // 4 3 2 1
			// 0 1 2 3
			// 0 1 2
			// 0 1
			// 0
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{24, 21, 6, 1, 15}
	
	sorArr := bubbleSort(arr)
	
	fmt.Println(sorArr)
	
}
