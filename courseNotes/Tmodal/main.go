package main

import (
	"fmt"
	"really"
)

const url  = "https://baidu.com"

type Retriever interface {
	Get(url string) string
}

func download(retriever Retriever) string {
	return retriever.Get(url)
}

func post(poster Poster) string {
	return poster.Post(url)
}

type Poster interface {
	Post(url string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}
func session(s RetrieverPoster) {
	s.Post(url)
	s.Get(url)
}

func main()  {
	var r  Retriever
	var p Poster
	var rp RetrieverPoster

	r = really.Retriever{}
	p = really.Poster{}
	rp = really.RetrieverPoster{}

	fmt.Println(rp.Get(url))
	fmt.Println(post(p))
	fmt.Println(download(r))
}