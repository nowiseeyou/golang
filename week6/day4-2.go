// 生成随机文本 ： 马尔可夫链算法 http://docscn.studygolang.com/doc/codewalk/markov/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Prefix 为拥有一个或多个单词的链 马尔可夫链的 前缀
type Prefix []string

// String 将 Prefix 作为一个 （用作映射键的） 字符串返回
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift 从Prefix 中移除第一个单词并追加上给定的单词
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

type Chain struct {
	chain       map[string][] string
	prefixLen   int
}

// NewChain 返回一个拥有 prefixLen 个单词前缀的 Chain
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

// Build 从提供的 Reader 中读取文本，并将它解析为存储了 前缀和后缀的 Chain
func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := make(Prefix, c.prefixLen)
	for {
		var s string
		if _,err := fmt.Fscan(br,&s); err != nil {
			break
		}
		key := p.String()
		c.chain[key] = append(c.chain[key],s)
		p.Shift(s)
	}
}

// Generate 返回一个从 Chain 生成的， 最多有 n 个单词的字符串
func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words,next)
		p.Shift(next)
	}
	return strings.Join(words, " ")
}

func main() {
	// 寄存命令行标记。
	numWords := flag.Int("words", 100, "maximum number of words to print")
	prefixLen := flag.Int("prefix", 2, "prefix length in words")
	
	flag.Parse()                     // 解析命令行标记。
	rand.Seed(time.Now().UnixNano()) // 设置随机数生成器的种子。
	
	c := NewChain(*prefixLen)     // 初始化一个新的 Chain。
	c.Build(os.Stdin)             // 从标准输入中构建链。
	text := c.Generate(*numWords) // 生成文本。
	fmt.Println(text)             // 将文本写入标准输出。
}
