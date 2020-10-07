package main

import (
	"encoding/json"
	"fmt"
)

// 仅当代码中不存在静态类型转换时才能这种声明
var _ json.Marshaler = (*json.RawMessage)(nil)

func main(){
	var v interface{}
	// m,ok := v.(json.Marshaler)
	if _,ok := v.(json.Marshaler); ok {
		fmt.Printf("value %v of type %T implements json.Marshaler \n", v, v)
	}
}