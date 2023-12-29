package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
}

var catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog map[int]Book) []Book {
	books := make([]Book, 0, len(catalog))
	for _, book := range catalog {
		books = append(books, book)
	}
	return books
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	book, ok := catalog[id]
	if !ok {
		return Book{}, fmt.Errorf("no book with id %d found", id)
	}
	return book, nil
}

func (b Book) NetPriceCents() int {
	return (b.PriceCents * (100 - b.DiscountPercent)) / 100
}
