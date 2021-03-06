package main

import (
	"fmt"
	"io"
	"reflect"
	"unicode"
)

func functionOfSomeType(v interface{}) interface{} {
	return reflect.TypeOf(v)
}

func nextInt(r []rune, i int) (value, nexPos int) {
	for ; i < len(r) && !unicode.IsDigit(r[i]); i++ {
	
	}
	x := 0
	for ; i < len(r) && !unicode.IsDigit(r[i]); i++ {
		x = x*10 + int(r[i]) - 0
	}
	
	return
}

func ReadFull(r io.Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
		
	}
	return
}

func main() {
	r := []rune("23kill")
	for i := 0; i < len(r); i++ {
		x, i := nextInt(r, i)
		fmt.Println(x)
		fmt.Println(i)
	}
	
	return
	// 判断接口变量类型
	var t interface{}
	t = functionOfSomeType(t)
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T \n", t) // %T 输出 t 是什么类型
	case bool:
		fmt.Printf("boolean %t \n", t)
	case int:
		fmt.Printf("integer %d \n", t)
	case *bool:
		fmt.Printf("pointer to bool %t \n", *t)
	case *int:
		fmt.Printf("point to int %d \n", *t)
	}
	
}
