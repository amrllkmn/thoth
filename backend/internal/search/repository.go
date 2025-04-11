package search

import (
	"github.com/amrllkmn/thoth/backend/internal/utils"
	"gorm.io/gorm"
)

type SQLiteBookRepository struct {
	db *gorm.DB
}

func (r *SQLiteBookRepository) FindAll() ([]utils.Book, error) {
	var books []utils.Book
	if err := r.db.Limit(20).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil

}
func (r *SQLiteBookRepository) FindByQuery(query string) {

}
func (r *SQLiteBookRepository) FindByID(id uint) {

}

func NewSQLiteBookRepository(db *gorm.DB) utils.BookRepository {
	return &SQLiteBookRepository{db: db}
}
