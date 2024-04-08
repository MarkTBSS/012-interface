package main

import (
	"fmt"

	models "github.com/MarkTBSS/012-interface/book"
)

type IObject interface {
	Getter() models.Book
	Setter(name string)
}

func NewBook(id int, name, author string) IObject {
	return &models.Book{
		ID:     id,
		Name:   name,
		Author: author,
	}
}

func main() {
	book1 := models.Book{
		ID:     1,
		Name:   "book1",
		Author: "bookAuthor1",
	}
	fmt.Println(book1)

	b := book1.Getter()
	fmt.Println(b)

	book1.Setter("Name Changed")
	fmt.Println(book1)

	book2 := NewBook(2, "book2", "bookAuthor2")
	fmt.Println(book2)
}
