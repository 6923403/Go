package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	var bok Books
	var bpoint *Books
	bok.title = "www"
	bok.author = "vcvc"

	bpoint = &bok
	fmt.Printf("title: %s, author: %s \n", bpoint.title, bpoint.author)
}

func test1() {
	fmt.Println(Books{"a", "bc", "def", 10})
	var T Books
	T.title = "qqq"
	fmt.Printf("%s \n", T.title)

	cout(T)
}

func cout(T Books) {
	fmt.Printf("%s \n", T.title)
}