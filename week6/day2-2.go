package main

import (
	"fmt"
	"log"
)

// 内嵌

type Reader interface {
	Reade(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter 接口结合了 Reader 和 Writer 接口
// type ReadWriter interface {
// 	Reader
// 	Writer
// }

// ReadWriter2 存储指向了 Reader 和 Writer 的指针
// 它实现了 io.ReadWriter
type ReadWriter struct {
	reader *Reader
	writer *Writer
}

// 为了提升该字段的方法并满足 io 接口，我们同样需要提供转发的方法

func (rw *ReadWriter) Read(p []byte) (n int, err error) {
	return rw.Read(p)
}


type Job struct {
	Command string
	*log.Logger
}

func NewJob (command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}

func (job *Job) Logf(format string, args ...interface{}) {
	job.Logf("%q: %s",job.Command,fmt.Sprintf(format,args...))
}








