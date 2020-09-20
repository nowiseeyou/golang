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

//值接收者
func (node treeNode) print(){
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

//中序遍历:先遍历左子树再遍历根，再遍历根节点再遍历右子树
func (node *treeNode) traverse(){
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main(){

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
	//root.traverse()


	root.print()
	fmt.Println(root.left)
	fmt.Println(root.right)
	fmt.Println(root.left.right)
	fmt.Println(root.right.left)

	//中序遍历
	fmt.Println("中序遍历")
	root.traverse()
	return

	//nodes
	nodes := []treeNode{
		{value:3},
		{},
		{6,nil,&root},
	}


	fmt.Println(nodes)

	//nil
	var pRoot *treeNode
	//return
	pRoot.setValue(200)
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print()

}

