package main

import (
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download(retriever Retriever) string {
	return retriever.Get("http://www.baidu.com")
}

func main()  {
	var r  Retriever
	fmt.Println(download(r))
}