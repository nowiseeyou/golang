package queue

type Queue []int

//push()方法可以在数组的末属添加一个或多个元素
//shift()方法把数组中的第一个元素删除
//unshift()方法可以在数组的前端添加一个或多个元素
//pop()方法把数组中的最后一个元素删除


func (q *Queue) Push(v int){
	*q = append(*q,v)
}

//删除第一个元素    并返回被删除元素值
func (q *Queue) Pop() int{
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
