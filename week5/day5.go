package main

import (
	"flag"
	"fmt"
	"log"
)

// 常量
type ByteSize float64

// 通过赋予空白标识符来忽略第一个值 (ignore first value by assigning to blank identifier
const (
	_   = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) Strings() string {
	switch  {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b <= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b/KB)
}

// init

func init() {
	if user == "" {
		log.Fatal("$USER not set")
	}
	
	if home == "" {
		home = "/home/" + user
	}
	
	if gopath == "" {
		gopath = home + "/go"
	}
	
	// gopath   可通过命令执行中在 --gopath 标记覆盖掉。
	flag.StringVar(&gopath, "gopath", gopath,"override default GOPATH")
}

func main() {

	xe13 := ByteSize.Strings(1e13)
	fmt.Println(xe13)
	
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(YB)
	fmt.Println(ZB)
	
}