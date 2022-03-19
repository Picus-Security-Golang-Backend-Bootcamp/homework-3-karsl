package author

import (
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
	authors, err := readFromCsv("author.csv")
	if err != nil {
		return err
	}

	for _, c := range authors {
		r.db.Create(&c)
	}

	return nil
}
