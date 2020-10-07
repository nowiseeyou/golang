package main

import (
	"fmt"
	"net/http"
	"time"
)

// 并发 不要通过共享内存来通信，而是通过通信来共享内存

// Go 程
// go list.Sort() // 并发运行 list.Sort,无需等它结束

func Announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}() // 注意括号 - 必须调用该函数
}

// 信道
// 信道 与映射一样，也需要通过make 来分配内存。其结果值充当了对底层数据结构的引用。
// 若提供了一个可选的整数形参，他就会为该信道设置缓冲区大小，默认值是零，表示不带缓冲的或同步的信道

// cj := make(chan int)            // 整数类型的无缓冲信道
// cj := make(chan int, 0)         // 整数类型的无缓冲信道
// cs := make(chan *os.File, 100)  // 指向文件指针的带缓冲信道

type listFace interface {
	Sort()
}

type List struct {
}

func (list List) Sort() {
	// TODO :
}

// 在 Go 程中启动排序。当它完成后，在信道上发送信号。
func xx() {
	c := make(chan int)     // 分配一个信道
	var list List
	go func() {
		list.Sort()
		c <- 1              // 发送信号，什么值无所谓
	}()
	// TODO :
	doSomethingForAWhile()
	<-c // 等待排序结束，丢弃发来的值
}

// 接受者在收到数据前会一直阻塞。若信道是不带缓冲的，那么在接受者收到值前，发送者会一直阻塞;
// 若信道是带缓冲的，则发送者仅在值被复制到缓冲区前阻塞;
// 若缓冲区已满，发送者会一直等待直到某个接受者取出一个值为止

// 信道缓冲区的容量决定了同时调用 process 的数量上限，因此我们在初始化时首先庶填充至它的容量上限
var MaxOutstanding int

var sem = make(chan int, MaxOutstanding)

func process(r interface{}) error {
	// TODO :
	return nil
}

func handle(r *http.Request) {
	sem <- 1   // 等待活动队列清空
	process(r) // 可能需要很长时间
	<-sem      // 完成：使下一个请求可以运行
}

// 设计问题 ： 尽管只有 MaxOutstanding 个 Go程 能同时运行 ，
// 但 Serve 还是为每个进入的请求都创建了新的Go 程。
// 若请求来的很快，该程序就会无线被消耗资源。
func Serve(queue chan *http.Request) {
	for {
		req := <-queue
		go handle(req) // 无需等待 handle 结束
	}
}


func Serve2(queue chan *http.Request) {
	for req := range queue {
		sem <- 1
		go func() {
			process(req) // BUG
			<-sem
		}()
	}
}

// BUG 出现在 GO 的 for 循环中，使循环变量在每次的迭代时会被重用，因此 req 变量会在 所有 Go程间共享，
// 我们需要确保 req 对于每个 Go程来说都是唯一的 ，有一种方法能够做到，就是将 req 的值作为实参传入到Go程的闭包中
func Serve3(queue chan *http.Request) {
	for req := range queue {
		sem <- 1
		go func (req *http.Request){
			process(req)
			<-sem
		}(req)
	}
}

func Serve4(queue chan *http.Request) {
	for req := range queue {
		req := req  // 为该 GO程创建 req 的新实例
		sem <- 1
		go func() {
			process(req)
			<-sem
		}()
	}
}

// 启动固定数量handle 一起从请求信道中读取数据，Go程限制了同时调用 process 的数量。
// Serve 同样会接受一个通知退出信道，在启动所有 Go程后，他将阻塞并暂停从信道中接收消息

func handle(queue chan *http.Request) {
	for r := range queue {
		process (r)
	}
}


func Serve5(clientRequests chan *http.Request, quit chan bool) {
	//  启动处理程序
	for i := 0; i < MaxOutstanding; i++ {
		go handle(clientRequests)
	}
	
	<-quit   // 等待通知退出
}

// 信道中的 信道
type Request struct {
	args        []int
	f           func([]int) int
	resultChan  chan int
}

func sum(a []int) (s int) {
	for _,v := range a {
			s += v
		}
	return
}

func xxx()  {
	request := &Request{[]int{3,4,5},sum,make(chan int)}
	// var clientRequest chan *http.Request
	// 发送请求
	clientRequest <- request
	
	// 等待回应
	fmt.Printf("Answer : %d \n", <-request.resultChan)
}

func handle2(queue chan *Request) {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}









