package main

import "fmt"


type treeNode struct {
	value int
	left,right *treeNode
}

//工厂函数
func createNode(value int) *treeNode {
	return &treeNode{value:value}
}

func main(){
	fmt.Println()
	
	root := treeNode{
		value: 3,
		left:  nil,
		right: nil,
	}

	root.left = &treeNode{}
	root.left.right = createNode(2)
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode)

	nodes := []treeNode{
		{},
		{},
		{6,nil,&root},
	}
	fmt.Println(nodes)
	fmt.Println(root)
}

