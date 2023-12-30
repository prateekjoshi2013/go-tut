package bookstore

import (
	"errors"
	"fmt"
)

const (
	CategoryAutobiography     = "Autobiography"
	CategoryLargePrintRomance = "Large Print Romance"
	CategoryParticlePhysics   = "Particle Physics"
)

var validCategory = map[string]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
	category        string
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
	if !validCategory[category] {
		return errors.New("invalid category")
	}
	b.category = category
	return nil
}
