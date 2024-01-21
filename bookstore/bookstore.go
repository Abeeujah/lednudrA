package bookstore

import (
	"errors"
	"fmt"
)

const (
	CategoryAutobiography = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

type Category int

// Book represents information about a book.
type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        Category
}

type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("No copies left")
	}
	b.Copies--
	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	// return books
	books := []Book{}

	for _, book := range c {
		books = append(books, book)
	}

	return books
}

func (c Catalog) GetBook(id int) (Book, error) {
	book, ok := c[id]
	if ok {
		return book, nil
	}
	return Book{}, fmt.Errorf("Error retrieving book, Invalid ID: %d specified", id)
}

func (book Book) NetPriceCents() int {
	saving := book.PriceCents * book.DiscountPercent / 100
	return book.PriceCents - saving
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("This is an Invalid amount, Price %d cannot be a negative value", price)
	}
	b.PriceCents = price
	return nil
}

func (b *Book) SetCategory(category Category) error {
	var validCategory = map[Category]bool{
		CategoryAutobiography:     true,
		CategoryLargePrintRomance: true,
		CategoryParticlePhysics:   true,
	}

	if !validCategory[category] {
		return fmt.Errorf("Specified an invalid book category: %q", category)
	}

	b.category = Category(category)
	return nil
}

func (b Book) Category() Category {
	return b.category
}
