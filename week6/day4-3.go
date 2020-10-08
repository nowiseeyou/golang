// 通过通信共享内存  http://docscn.studygolang.com/doc/codewalk/sharemem/

package main

import (
	"log"
	"net/http"
	"time"
)

const (
	numPollers     = 2                // Poller Go 程的启动数
	pollInterval   = 60 * time.Second // 轮询每一个 URL 频率
	statusInterval = 10 * time.Second // 将状态记录到标准输出的频率
	errTimeout     = 10 * time.Second // 回退超时的错误
)

var urls = []string{
	"http://www.google.com",
	"http://golang.org",
	"http://blog.golang.org",
	"http://www.test.net",
	"http://www.test.xyz",
}

type State struct {
	url    string
	status string
}

// logState 打印
func logState(s map[string]string) {
	log.Println("Current state: ",s)
	for k, v := range s {
		log.Printf(" %s %s ", k, v)
	}
}

// StateMonitor 维护了一个映射， 它存储了URL被轮询的状态，并每隔 updateInterval 纳秒
// 打印出其当前的状态，它向资源状态的接受者返回一个 chan State

func StateMonitor(updateInterval time.Duration) chan<- State {
	updates := make(chan State)
	urlStatus := make(map[string]string)
	ticker := time.NewTicker(updateInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				logState(urlStatus)
			case s := <-updates:
				urlStatus[s.url] = s.status
				
			}
		}
	}()
	return updates
}

// Response 表示一个被此程序轮询的 HTTP URL
type Resource struct {
	url      string
	errCount int
}

// Poll 为 url 执行一个 HTTP HEAD 请求，并返回 HTTP 的状态字符或一个错误字符串
func (r *Resource) Poll() string {
	resp,err := http.Head(r.url)
	if err != nil {
		log.Print("Error", r.url, err)
		r.errCount++
		return err.Error()
	}
	r.errCount = 0
	return resp.Status
}

// Sleep 在将 Resource 发送到 done 之前休眠一段适当时间（取决于错误状态
func (r *Resource) Sleep(done chan <- *Resource) {
	time.Sleep(pollInterval + errTimeout * time.Duration(r.errCount))
}

func Poller(in <- chan *Resource, out chan <- *Resource, status chan<- State) {
	for r := range in {
		s := r.Poll()
		status <- State{r.url,s}
		out <- r
	}
}

func main() {
	// 创建我们的输入和输出的信道
	pending,complete := make(chan *Resource), make(chan *Resource)
	
	// 启动 StateMonitor
	status := StateMonitor(statusInterval)
	
	// 启动一些 poller Go程
	for i := 0; i < numPollers; i++ {
		go Poller(pending, complete, status)
	}
	
	// 将一些 Resource 发送至 pending 序列
	go func() {
		for _,url := range urls {
			pending <- &Resource{url: url}
		}
	}()
	
	for r := range complete{
		go r.Sleep(pending)
	}
}
