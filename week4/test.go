package main

import (
	"fmt"
	"reflect"
	"unicode"
)

func functionOfSomeType( v interface{}) interface{}{
	return reflect.TypeOf( v )
}

func newInt(r []rune , i int) (int,int) {
	for ; i < len(r) && !unicode.IsDigit(r[i]); i++ {
	
	}
	x := 0
	for ; i < len(r) && !unicode.IsDigit(r[i]); i++ {
		x = x * 10  +  int(r[i]) - 0
	}
	
	return x,i
}

func main() {
	// 判断接口变量类型
	var  t interface{}
	t = functionOfSomeType(t)
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T \n", t)        // %T 输出 t 是什么类型
	case bool:
		fmt.Printf("boolean %t \n", t)
	case int:
		fmt.Printf("integer %d \n",t)
	case *bool:
		fmt.Printf("pointer to bool %t \n", *t)
	case *int:
		fmt.Printf("point to int %d \n", *t)
	}
	
}
