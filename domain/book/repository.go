package book

import (
	"errors"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-karsl/infrastructure/data"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r BookRepository) Migration() error {
	err := r.db.AutoMigrate(&Book{})
	if err != nil {
		return err
	}

	return nil
}

// InsertSampleData reads data from book.csv and writes them to table book
func (r BookRepository) InsertSampleData() error {
	lines, err := data.GetCellsFromCSV("book.csv")
	if err != nil {
		return err
	}

	books, err := linesToBook(lines)
	if err != nil {
		return nil
	}

	for _, c := range books {
		r.db.FirstOrCreate(&c, Book{Name: c.Name})
	}

	return nil
}

// Search returns all books that match with term in the bookshelf, by checking if it matches with given term by
// considering Name, Author.Name, StockCode fields of Book.
func (r BookRepository) Search(term string) []Book {
	var foundBooks []Book
	term = strings.ToLower(term)
	r.db.Preload(clause.Associations).Joins("JOIN Author on Author.id = Book.author_id").
		Where("lower(Book.name) LIKE ?", "%"+term+"%").
		Or("lower(Author.name) LIKE ?", "%"+term+"%").
		Or("lower(Book.stock_code) LIKE ?", "%"+term+"%").
		Find(&foundBooks)

	return foundBooks
}

// List returns all books in the bookshelf.
func (r BookRepository) List() []Book {
	var books []Book
	r.db.Find(&books)
	return books
}

// Buy find the book with given id and buys it.
func (r BookRepository) Buy(bookId, quantity int) error {
	if quantity <= 0 {
		return errors.New("invalid quantity")
	}

	foundBook, err := r.findBookById(bookId)
	if err != nil {
		return err
	}

	err = foundBook.buy(quantity)
	if err != nil {
		return err
	}

	r.db.Save(foundBook)

	return nil
}

// Delete deletes the given book
func (r BookRepository) Delete(bookID int) error {
	result := r.db.Delete(&Book{}, bookID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// findBookById returns the first book with given id in books.
func (r BookRepository) findBookById(bookId int) (Book, error) {
	var book Book
	result := r.db.First(&book, bookId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return Book{}, errors.New("no such book found")
	}

	return book, nil
}
