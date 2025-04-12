package search

import (
	"errors"

	"github.com/amrllkmn/thoth/backend/internal/database"
	"github.com/amrllkmn/thoth/backend/internal/utils"
	"gorm.io/gorm"
)

type SQLiteBookRepository struct {
	db *gorm.DB
}

func (r *SQLiteBookRepository) FindAll(page, limit int) ([]utils.Book, error) {
	var books []utils.Book
	if err := r.db.Scopes(database.Paginate(page, limit)).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil

}
func (r *SQLiteBookRepository) FindByQuery(query string) ([]utils.Book, error) {
	var books []utils.Book
	if err := r.db.Where("title LIKE ?", "%"+query+"%").Or("authors LIKE ?", "%"+query+"%").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
func (r *SQLiteBookRepository) FindByID(isbn string) (*utils.Book, error) {
	var book utils.Book
	results := r.db.Where("isbn13 = ? OR isbn10 = ?", isbn, isbn).First(&book)

	if errors.Is(results.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if results.Error != nil {
		return nil, results.Error
	}

	return &book, nil
}
func NewSQLiteBookRepository(db *gorm.DB) utils.BookRepository {
	return &SQLiteBookRepository{db: db}
}
