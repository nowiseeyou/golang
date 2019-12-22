package main

import (
	"./really"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download(retriever Retriever) string {
	return retriever.Get("https://baidu.com")
}

func main()  {
	var r  Retriever
	r = really.Retriever{}
	fmt.Println(download(r))
}