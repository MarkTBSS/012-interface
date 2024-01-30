package main

import (
	"fmt"

	"github.com/MarkTBSS/go-interface/models"
)

type IprivateStructBook interface {
	Getter() models.PublicStructBook
	Setter(name string)
}

type privateStructBook struct {
	id     int
	name   string
	author string
}

func NewBook(id int, name, author string) IprivateStructBook {
	return &privateStructBook{
		id:     id,
		name:   name,
		author: author,
	}
}

// Getter is a method for privateStructBook that returns a PublicStructBook
func (b *privateStructBook) Getter() models.PublicStructBook {
	return models.PublicStructBook{
		Id:     b.id,
		Name:   b.name,
		Author: b.author,
	}
}

func (b *privateStructBook) Setter(name string) {
	b.name = name
}

func main() {
	book1 := privateStructBook{
		id:     1,
		name:   "book1",
		author: "bookAuthor1",
	}
	fmt.Println(book1)

	// Use the Getter method to convert privateStructBook to PublicStructBook
	publicBook := book1.Getter()
	fmt.Println(publicBook)

	// Use the Setter method to modify the 'name' field
	book1.Setter("Name Changed")
	fmt.Println(book1)

	book2 := NewBook(2, "book2", "bookAuthor2")
	fmt.Println(book2)
}
