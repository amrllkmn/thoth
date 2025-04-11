package search

import (
	"github.com/amrllkmn/thoth/backend/internal/utils"
	"gorm.io/gorm"
)

type SQLiteBookRepository struct {
	db *gorm.DB
}

func (r *SQLiteBookRepository) FindAll() {

}
func (r *SQLiteBookRepository) FindByQuery(query string) {

}
func (r *SQLiteBookRepository) FindByID(id uint) {

}

func NewSQLiteBookRepository(db *gorm.DB) utils.BookRepository {
	return &SQLiteBookRepository{db: db}
}
