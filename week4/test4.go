package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"sync"
)

type SyncedBuffer struct {
	lock sync.Mutex
	buffer bytes.Buffer
}

// func NewFile (fd int, name string ) *File{
// 	if fd < 0 {
// 		return nil
// 	}
//
// 	f := new (File)
// 	f.fd = fd
// 	f.name = name
// 	f.dirinfo = nil
// 	f.nepipe = 0
// 	return f
//
// }

func NewFile(fd int, name string ) *File {
	if fd < 0 {
		return nil
	}
	
	// f := File{fd, name, nil, 0}
	// return &File{fd, name, nil, 0}
	return &File{fd: fd, name: name}
}

func main(){
	p := new (SyncedBuffer) // type *SyncedBuffer
	var v SyncedBuffer      // type  SyncedBuffer
	fmt.Print(p)
	fmt.Print(v)
	
	// 创建数组 切片
	a := [...]string  {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string  {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := [...]string  {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
}


