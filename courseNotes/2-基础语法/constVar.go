package main

import "fmt"

const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

func main()  {
	fmt.Println(b,kb,mb,gb,tb,pb)
}