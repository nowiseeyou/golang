package main

import "fmt"

// 冒泡算法

func bubbleSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len; i++ {  // 1-4 2-3 3-2 2-1 1
		for j := 0; j < len-1-i; j++ {  // 4 3 2 1
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


// 快速排序
func selectionSort(arr []int ) []int {
	len := len(arr)
	for i := 0; i < len - 1; i++{
		min := i
		for j := i+1; j<len; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	
	return arr
}

func main() {
	arr := []int{24, 21, 6, 1, 15}
	
	bubbleSortArr := bubbleSort(arr)
	selectionSort := selectionSort(arr)
	
	fmt.Println(bubbleSortArr)
	fmt.Println(selectionSort)
	
}
