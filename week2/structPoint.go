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
	var book2 Books
	
	book1.title = "golang"
	book1.author = "Google"
	book1.subject = "编程语言"
	book1.book_id = 12334
	
	book2.title = "php"
	book2.author = "AndyLee"
	book2.subject = "编程语言"
	book2.book_id = 12443
	
	printBook(&book1)
	
	printBook(&book2)
	
}

func printBook(book *Books) {
	fmt.Println(book.title)
	fmt.Println(book.author)
	fmt.Println(book.subject)
	fmt.Println(book.book_id)
}
