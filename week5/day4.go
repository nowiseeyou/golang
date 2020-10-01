package main

import (
	"fmt"
	"log"
	"os"
)

type T struct {
	a int
	b float64
	c string
}

func (t *T) String() string{
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}


type MyString string

func (m MyString) String2() string {
	// 注意： 请勿通过调用 Sprintf 来构造 String 方法，因为它会无限递归你的的 String2 方法
	// return fmt.Sprintf("Mystring = %s", m) // 错误： 会无限递归
	return fmt.Sprintf("Mystring = %s",string(m))
}

type interfaceName interface {
	a()
}
// ...interface{} 类型，这样格式的后面就能出现任意数量，任意类型的形参了
// Println 通过 fmt.Println 的方式将日志打印到标准记录器
func Println( v ...interface{}){
	log.Output(2,fmt.Sprintln(v...)) // Output 接收形参(int, string)
}

// ...形参可指定具体的类型 *** ^ ***
func Min(a ...int) int {
	min := int(^uint(0) >> 1) // 最大的 int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

// 打印
func main() {
	fmt.Printf("Hello %d \n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))
	
	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x ; %d %x \n", x, x, int64(x), int64(x))
	
	s := "nice too meet you"
	
	t := &T{7, -2.35, "ABC\tdef"}
	fmt.Printf("%v \n", t)
	fmt.Printf("%+v \n", t)
	fmt.Printf("%#v \n", t)
	
	// fmt.Printf %T 打印类型
	fmt.Printf("%T \n", s)
	
	rst := t.String()
	fmt.Println(rst)
	
	itf := interfaceName.a
	Println(itf)
	
}
