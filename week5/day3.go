package main

import (
	"fmt"
	"log"
)

// 映射 集合
var timeZone = map[string]int {
	"UTC" :  0 * 60 * 60,
	"EST" : -5 * 60 * 60,
	"CST" : -6 * 60 * 60,
	"MST" : -7 * 60 * 60,
	"PST" : -8 * 60 * 60,
}

// 多重赋值
var seconds int
var ok bool

func offset(tz string) int{
	if seconds, ok := timeZone[tz];ok {
		return seconds
	}
	
	log.Println("unknown time zone : ",tz)
	return 0
}

func main()  {

	fmt.Println(timeZone)
	
	attended := map[string]bool{
		"Anne" : true,
		"Jone" : true,
	}
	person := "Alice"
	if attended[person] == false{ // 若某人不在此映射中 则为false
		fmt.Println(person," 正在开会")
	}
	
	delete(timeZone,"PST")
	
	offset := offset("PRC")
	fmt.Println(offset)
}