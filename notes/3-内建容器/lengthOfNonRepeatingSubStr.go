package main

import (
	"encoding/json"
	"fmt"
)

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
	//lengthOfNonRepeatingSubStr("aaacbaa")
	//return
	fmt.Println([]byte("aahello"))
	fmt.Println([]rune("233"))
	fmt.Println(len("881DB6CCE417CBD0D30D0D29EC3559E2"))
	fmt.Println(lengthOfNonRepeatingSubStr("aaaabcddabcafeaaa"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二一三三儿一二1R一一"))

	return
	type profile struct {
		name string
		age int
		gender int
		hobby string
	}

	theShy := profile{
		name:"姜承碌",
		age:19,
		gender:1,
		hobby:"发育",
	}
	json,err := json.Marshal(theShy)

	fmt.Println(theShy)
	fmt.Println(theShy.name)
	fmt.Println(theShy.age)
	fmt.Println(theShy.gender)
	fmt.Println(theShy.hobby)

	fmt.Println(string(json),err)
}
