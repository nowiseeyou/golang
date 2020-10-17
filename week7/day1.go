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

// 插入排序
func insertionSort(arr []int) []int {
	for i := range arr  {
		preIndex := i
		current := arr[i]       // arr[1]
		for preIndex > 0 && arr[preIndex-1] > current  {        // arr[0] > arr[1]
			arr[preIndex] = arr[preIndex-1]         // arr[1] = arr[0]
			preIndex--
		}
		arr[preIndex] = current     // arr[0] = arr[1]
	}
	return arr
}

func main() {
	arr := []int{24, 21, 6, 1, 15}
	
	bubbleSortArr := bubbleSort(arr)        // 冒泡排序
	selectionSort := selectionSort(arr)     // 快速排序
	insertionSort := insertionSort(arr)     // 插入排序
	
	fmt.Println(bubbleSortArr)
	fmt.Println(selectionSort)
	fmt.Println(insertionSort)
	
}
