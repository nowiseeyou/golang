package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var book1 Books
	
	// book1 描述
	book1.title = "朝花夕拾"
	book1.author = "鲁迅"
	book1.subject = "文学"
	book1.book_id = 12138
	
	fmt.Println(book1)
}
