package main

import (
	"fmt"
	"sort"
)

// 接口

type Sequence []int

// Methods required by sort.Interface
// sort.Interface 所需方法

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i],s[j] = s[j],s[i]
}

// Method for printing - sorts the elements before printing .
// 用于打印的方法 - 在打印前对元素进行排序

func (s Sequence) String() string {
	sort.Sort(s)
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	
	return str + "]"
}

func main(){
	Se := Sequence{5,2,9,4,6,8,4}
	sortString := Se.String()
	fmt.Println(sortString)
}