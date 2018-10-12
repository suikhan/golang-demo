package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main(){
	var Book1, Book2 Books

	Book1.title="Go 语言"
	Book1.author="www.runoob.com"
	Book1.subject = "GO语言教程"
	Book1.book_id = 6495407

	Book1.title="Go 语言"
	Book1.author="www.runoob.com"
	Book1.subject = "GO语言教程"
	Book1.book_id = 6495700

	fmt.Printf("Book1 title: %s\n", Book1.title)
	fmt.Printf("Book1 author: %s\n", Book1.author)
	fmt.Printf("Book1 subject: %s\n", Book1.subject)
	fmt.Printf("Book1 book_id: %d\n", Book1.book_id)

	
	fmt.Printf("Book2 title: %s\n", Book2.title)
	fmt.Printf("Book2 author: %s\n", Book2.author)
	fmt.Printf("Book2 subject: %s\n", Book2.subject)
	fmt.Printf("Book2 book_id: %d\n", Book2.book_id)
}