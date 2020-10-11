package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

// 切片 切片基于数组构建

// 切片写法 []T, T是切片元素的类型，和数组不同，切片类型并没有固定的长度

// letters := []string{"a","b","c","d"}

// func make([]T,len,cap) []T
// 其中T代表被创建的切片元素的类型。函数 make 接收一个类型、长度和可选容量参数。
// 调用 make 时，内部会分配一个数组，然后返回数组对应的切片
// 当容量参数被忽略，它默认为指定的长度


// 切片的内幕 ：
// 一个切片是一个数组片段的描述。它包含了指向数组的指针（ptr），片段的长度（len），和容量（cap）（片段的最大长度）
// 通过一个新切片修改元素 会 影响到原始切片的对应元素


// 切片的生长 （copy and append 函数）

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate  (如有必要，请重新分配
		// allocate double what's needed, for future growth.
		newSlice := make([]byte ,(n+1)*2)
		copy(newSlice,slice)
		slice = newSlice
	}
	
	slice = slice[0:n]
	copy(slice[m:n],data)
	return slice
}


// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter (s []int, fn func(int) bool) []int{
	var p []int   // == nil
	for _,v := range s {
		if fn(v) {
			p = append(p,v)
		}
	}
	
	return p
}


// ===========================陷阱=================================

// 切片操作并不会复制底层的数组。整个数组将被保存在内存中，直到它不在被引用。
// 有时候可能会因为 一个小的内存引用 导致 保存所有的数据


var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b,_ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}


// 这段代码的行为和描述类似，返回的 []byte 指向保存整个文件的数组。
// 因为切片引用了原始数组，导致GC不能释放数组的空间;
// 只用到少数几个字节却导致整个文件的内容都一直保存在内存里
// 要修复整个问题，可以将感兴趣的数据复制到一个新的切片中

func CopyDigits(filename string) []byte {
	b,_ := ioutil.ReadFile(filename)
	b  = digitRegexp.Find(b)
	c := make([]byte, len(b))
	
	copy(c,b)
	return c
}

// =================================================================



func main()  {
	letters := []string{"a","b","c","d"}
	fmt.Printf("切片letters : %s \n ", letters)
	
	var s []byte
	s = make([]byte, 5, 5)      // len(s) == 5   cap(s) == 5
	fmt.Printf("切片s : %s \n ", s)
	
	x1 := [3]string {"a","b","c"}
	s1 := x1[:]
	fmt.Println(s1)
	
	t := make([]byte, len(s), (cap(s) + 1) * 2)   // +1 in case cap(s) == 0
	for i := range s {
		t[i] = s[i]
	}
	s = t
	
	// copy
	var s2 []byte
	t2 := make([]byte, len(s2),(cap(s2) + 1) * 2)
	copy(t2,s2)
	s2 = t2
	
	fmt.Println(len(s2))
	
	p := []byte{2,3,5}
	p = AppendByte(p,7,11,13)
	
	fmt.Println(p)
	
	// append
	// func append(s []T,x ...T) []T
	a := make([]int, 1)
	// a == []int{0}
	a = append(a,1,2,3)
	// a == []int{0,1,2,3}
	
	// 如果将一个切片追加到另一个切片尾部，需要使用 ... 语法将第2个参数 展开为参数列表
	
	a1 := []string{"John", "Paul"}
	b1 := []string{"George","Ringo","Pete"}
	a1 = append(a1,b1...)
	// a == []string{"John","Paul","George","Ringo","Pete"}
	
	fmt.Println(a1)
	
}

