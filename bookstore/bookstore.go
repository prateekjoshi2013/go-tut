package bookstore

import (
	"errors"
	"fmt"
)

type Category string

func (c Category) ValidateCategory() error {
	switch c {
	case "Romance", "Western", "Quantum Mechanics":
		return nil
	default:
		return errors.New("not a valid category")
	}
}

func (c Category) String() string {
	return string(c)
}

type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
	category        Category
}

type Catalog map[int]Book

var catalog Catalog

func (b Book) Buy() (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	books := make([]Book, 0, len(c))
	for _, book := range c {
		books = append(books, book)
	}
	return books
}

func (c Catalog) GetBook(id int) (Book, error) {
	book, ok := c[id]
	if !ok {
		return Book{}, fmt.Errorf("no book with id %d found", id)
	}
	return book, nil
}

func (b Book) NetPriceCents() int {
	return (b.PriceCents * (100 - b.DiscountPercent)) / 100
}

func (b *Book) SetPriceCents(priceCents int) error {
	if priceCents < 0 {
		return errors.New("price cannot be negative")
	}
	b.PriceCents = priceCents
	return nil
}

func (b *Book) SetCategory(category string) error {
	c := Category(category)
	if err := c.ValidateCategory(); err != nil {
		return err
	}
	b.category = c
	return nil
}
