package main

import (
	"fmt"
	"net/http"
	"os"
)

type Phone interface {
	call()
}

type NokiaPhone struct {}

func (nokia NokiaPhone) call(){
	fmt.Println("I am Nokia,I can call you!")
}


type IPhone struct {}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone,I can call you!")
}

type Stringer interface {
	String() string
}

type strings struct {}

func (s strings ) String() string{
	return "I am Interface xi xi man sao de"
}

// 这是计数器模式CTR流的定义，它将块加密改为流加密，注意块加密的细节已被抽象化

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}

type Stream interface {
	XORKeyStream(dst, src []byte)
}

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// 记录某个页面被访问的次数
// type Counter int
//
// func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request){
// 	*ctr++
//
// 	fmt.Fprintf(w, "counter = %d \n", *ctr)
// }

// 每次浏览该信道都会发送一次提醒
// 可能需要 带缓冲的信道

type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprint(w, "notification sent")
}

func ArgServer() {
	fmt.Println(os.Args)
}

func main() {
	var phone Phone
	
	phone = new(NokiaPhone)
	phone.call()
	
	phone = new(IPhone)
	phone.call()
	
	var ser Stringer
	ser = new(strings)
	
	fmt.Println(ser.String())
}
