package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func fibonacci() intGen {
	a,b := 0,1
	return func() int {
		a,b = b, a+b
		return a
	}
}

type intGen func() int

//系统Reader
func (g intGen) Read(p []byte) (n int, err error){
	next := g()
	if next > 1000 {
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d \n",next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node)) {
		n.Print()
	}

	fmt.Println()
}

func(node *Node) TraverseFunc(f func(*Node)){
	if node == nil {
		return
	}

	node.left.TraverseFunc(f)
	f(node)
	node.right.TraverseFunc(f)
}



func main() {
	root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node)){
		nodeCount++
	}
	fmt.Println("Node count:",nodeCount)
	return
	f := fibonacci()
	printFileContents(f)

	return
	a  := adder()
	for i := 0 ;i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d \n",i,a(i) )
	}
}
