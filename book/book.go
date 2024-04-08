package models

type Book struct {
	ID     int
	Name   string
	Author string
}

func (b *Book) Getter() Book {
	return Book{
		ID:     b.ID,
		Name:   b.Name,
		Author: b.Author,
	}
}

func (b *Book) Setter(name string) {
	b.Name = name
}
