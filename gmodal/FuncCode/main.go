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

func Fibonacci() intGen {
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

func (node *treeNode) Traverse() {
	node.TraverseFunc(func(n *treeNode) {
		n.print()
	})

	fmt.Println()
}

func(node *treeNode) TraverseFunc(f func(*treeNode)){
	if node == nil {
		return
	}

	node.left.TraverseFunc(f)
	f(node)
	node.right.TraverseFunc(f)
}

type treeNode struct {
	value int
	left,right *treeNode
}

//工厂函数
func createNode(value int) *treeNode {
	return &treeNode{value:value}
}

//值接收者
func (node *treeNode) print(){
	fmt.Println(node.value," ")
}

//指针接收者
func (node *treeNode) setValue(value int){
	//nil 可以调用指针
	if node == nil {
		fmt.Println("setting value to nil node ignored" )
		return
	}
	node.value = value
}

func main() {
	root := treeNode{
		value: 3,
		left:  nil,
		right: nil,
	}

	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}

	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	root.right.left.setValue(4)
	root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *treeNode){
		nodeCount++
	})
	fmt.Println("Node count:",nodeCount)
	return
	f := Fibonacci()
	printFileContents(f)

	return
	a  := adder()
	for i := 0 ;i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d \n",i,a(i) )
	}
}
