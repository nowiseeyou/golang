package main

import (
	"fmt"
	"net/http"
	"os"
)

// HandlerFunc 类型是一个适配器， 它允许将普通函数用作 HTTP 处理程序
// 若 f 是个具有适当签名的函数， HandlerFunc(f) 就是调用 f 的处理程序对象

type HandlerFunc func(http.ResponseWriter,*http.Request)

// ServeHTTP call f(c, req)
func (f HandlerFunc) ServeHTTP(w http.ResponseWriter,r *http.Request){
	f(w,r)
}

// 实参服务器
func ArgServer(w http.ResponseWriter, req *http.Request){
	fmt.Fprintln(w,os.Args)
}

func main(){
	// 当有人访问  /args 页面时 安装到该页面的处理程序就有了值 ArgServer  和类型 HandlerFunc 。
	// HTTP服务器会以 ArgServer 为接受者，调用该类型的 ServerHTTP 方法，它会反过来调用 ArgServer(通过f(c, req)) ，
	// 接着
	http.Handle("/args", http.HandlerFunc(ArgServer))
}




