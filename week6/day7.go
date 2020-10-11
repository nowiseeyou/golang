package main

import (
	"fmt"
	"sort"
)

// interface 接口

type Person struct {
	Age        int
	Experience int
	Name       string
}

func (p Person) String() string {
	return fmt.Sprintf("%s : %d", p.Name, p.Age)
}

// ByAge implements sort. Interface for []Person based on the Age field.

type ByAge []Person // 自定义

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{31, 3, "Bob"},
		{42, 6, "John"},
		{17, 5, "Michael"},
		{26, 2, "Alice"},
	}
	
	fmt.Println(people)
	fmt.Println(ByAge(people))
	
	sort.Sort(ByAge(people))
	fmt.Println(people)
	
}
