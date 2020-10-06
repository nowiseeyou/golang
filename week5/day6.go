package main

import (
	"fmt"
	"sort"
)

type Sequence []int

// 用于打印的反法 - 在打印前对元素进行排序
func (s Sequence) String() string {
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

func main()  {
	se := Sequence {7,2,5,4,5}
	se.String()
	fmt.Println(se)
}