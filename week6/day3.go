package main

import (
	"bytes"
	"debug/macho"
)

// 并行化
type Vector []float64

// 将此操作 应用至  v[i], v[i+1] ... 直到 v[n-1]
func (v Vector) DoSome (i, n int,u Vector,c chan int){
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1      // 发信号表示这一块计算完成
}

// ==================2=================

const nCPU  = 4  // cpu 核心数

func (v Vector) DoAll (u Vector) {
	c := make(chan int,nCPU) // 缓冲区是可选的，但明显用上
	
	for i := 0; i<nCPU ; i++  {
		go v.DoSome(i * len(v) / nCPU, (i+1) * len(v) / nCPU, u, c)
	}
	
	// 排空信道
	for i := 0; i < nCPU; i++ {
		<- c   // 等待任务处理完成
	}
	
	// 一切完成
}

// 注意： 并发/并行
// 并发是用独立执行的组件构造的方法，而并行则是为了效率在多个CPU上平行的进行计算

// 可能泄露缓冲区

func load(b interface{}){}

var freeList = make(chan *bytes.Buffer, 100)
var serverChan = make(chan *bytes.Buffer)

func client(){
	for {
		var b *bytes.Buffer
		// 若缓冲区可用就用它，不可用就分配新的
		select {
		case b =<- freeList:
			// 获取一个，不做别的
			default:
			// 非空闲，因此分配一个新的
			b = new(bytes.Buffer)
		}
		load(b)             // 从网络中读取下一条消息
		serverChan <- b     // 发送至服务器
	}
}











