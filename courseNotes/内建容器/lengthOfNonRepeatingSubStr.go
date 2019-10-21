package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int{
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i,ch := range []rune(s) {
		if lastI,ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i - start +1 > maxLength {
			maxLength = i - start +1
		}

		lastOccurred[ch] = i
	}
	return maxLength
	//fmt.Println(lastOccurred)
}

func main(){
	fmt.Println(len("881DB6CCE417CBD0D30D0D29EC3559E2"))
	fmt.Println(lengthOfNonRepeatingSubStr("aaaabcddabcafeaaa"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二一三三儿一二1R一一"))
}
