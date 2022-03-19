package book

import (
	"errors"
	"fmt"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-karsl/domain/domain/author"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name          string
	StockCode     string
	ISBN          int
	NumberOfPages int
	Price         float64
	Quantity      int
	AuthorID      int
	Author        author.Author
	gorm.DeletedAt
}

func (book Book) BeforeDelete(tx *gorm.DB) error {
	fmt.Println("Deleting author: ", book.Name)
	return nil
}

func (book Book) String() string {
	return fmt.Sprintf("{Id: %d, Name: %s, StockCode: %s, ISBN: %d, NumberOfPages: %d, Price: %f, Quantity: %d, Author: %s}",
		book.ID, book.Name, book.StockCode, book.ISBN, book.NumberOfPages, book.Price, book.Quantity, book.Author)
}

// buy decreases stock count. A Book can't be bought if there is not enough stock or deleted already.
func (book *Book) buy(quantityToBuy int) error {
	if book.Quantity < quantityToBuy {
		return errors.New("there is not enough items in the stock")
	}

	book.Quantity -= quantityToBuy
	return nil
}
