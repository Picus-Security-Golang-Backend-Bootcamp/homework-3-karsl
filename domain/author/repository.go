package author

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-karsl/infrastructure/data"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (r *AuthorRepository) Migration() error {
	err := r.db.AutoMigrate(&Author{})
	if err != nil {
		return err
	}

	return nil
}

// InsertSampleData reads data from author.csv and writes them to table author
func (r *AuthorRepository) InsertSampleData() error {
	lines, err := data.GetCellsFromCSV("author.csv")
	if err != nil {
		return err
	}

	authors, err := linesToAuthors(lines)
	if err != nil {
		return err
	}

	for _, c := range authors {
		r.db.FirstOrCreate(&c, Author{
			Model: gorm.Model{ID: c.ID},
		})
	}

	return nil
}
