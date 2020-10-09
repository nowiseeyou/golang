package main

import (
	"fmt"
	"log"
	"regexp"
	"syscall"
)

// 错误

type error interface {
	Error() string
}

// PathError 记录一个错误以及产生该错误的路径和操作

type PathError struct {
	Op      string      // "open" , "unlink" 等等
	Path    string      // 相关联的文件
	Err     error       // 由系统调用返回
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + " :  " + e.Err.Error()
}

func veryClose(z, prevz float64) bool{
	// TODO:
	 x  := z + prevz
	 if x > 0 {return false}
	return true
}

// Panic 向调用者报告错误的一般方式就是将 error 作为额外的值返回

// 用牛顿法计算立方根的一个玩具实现
func CubeRoot(x float64) float64 {
	z := x/3    // 任意初始值
	for i := 0; i < 1e6; i++ {
		prevz := z
		z -= (z*z*z -x) / (3*z*z)
		if veryClose(z, prevz) {
			return z
		}
	}
	// 一百万次迭代并未收敛， 事情出错了
	panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}
var  work interface{}

// 恢复
func server(workChan <- chan *Work) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover();err != nil {
			log.Println("work failed:", err)
		}
	}()
	do(work)
}

// Error 是解析错误的类型，它满足 error 接口
type Error string

func (e Error) Error() string {
	return string(e)
}

// error 是*Regexp 的方法，通过一个 Error 来触发 Panic 来报告解析错误
func (regexp *regexp.Regexp) error(err string) {
	panic(Error(err))
}

// Compile 返回该正则表达式解析后的表示
func Compile(str string) (regexp *regexp.Regexp,err error){
	regexp = new(regexp.Regexp)
	// doParse will panic if there is a parse error
	defer func() {
		if e := recover(); e != nil {
			regexp = nil        // 清理返回值
			err  = e.(Error)    // 若它不是解析错误，将重新出发 Panic
		}
	}()
	return regexp.doParse(str), nil
}