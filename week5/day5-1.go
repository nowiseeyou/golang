package main

import "fmt"

// type ByteSlice []type
type ByteSlice []string

func (p *ByteSlice) Append(data []byte)  {
	// 主体和前面相同 没有 return
	slice := *p
	*p = slice
}

func (p *ByteSlice) Write(data []byte) (n int,err error) {
	slice := *p
	// 依旧和前面相同
	*p = slice
	return len(data),nil
}

func main(){
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days \n", 7)
}

