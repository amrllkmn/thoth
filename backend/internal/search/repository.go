package search

import "gorm.io/gorm"

type Book struct {
	ID            uint    `gorm:"primaryKey"`
	Isbn13        int64   `gorm:"unique" json:"isbn13"`
	Isbn10        string  `gorm:"unique" json:"isbn10"`
	Title         string  `json:"title"`
	Subtitle      string  `json:"subtitle"`
	Authors       string  `json:"authors"`
	Categories    string  `json:"categories"`
	Thumbnail     string  `json:"thumbnail"`
	Description   string  `json:"description"`
	PublishedYear int     `json:"published_year"`
	AverageRating float32 `json:"average_rating"`
	NumPages      int     `json:"num_pages"`
	RatingsCount  int     `json:"ratings_count"`
}

type BookRepository interface {
	FindAll()
	FindByQuery(query string)
	FindByID(id uint)
}

type SQLiteBookRepository struct {
	db *gorm.DB
}

func (r *SQLiteBookRepository) FindAll() {

}
func (r *SQLiteBookRepository) FindByQuery(query string) {

}
func (r *SQLiteBookRepository) FindByID(id uint) {

}

func NewSQLiteBookRepository(db *gorm.DB) BookRepository {
	return &SQLiteBookRepository{db: db}
}

type MeilisearchBookRepository struct {
}

func (r *MeilisearchBookRepository) FindAll()                 {}
func (r *MeilisearchBookRepository) FindByQuery(query string) {}
func (r *MeilisearchBookRepository) FindByID(id uint)         {}

func NewMeilisearchBookRepository() BookRepository {
	return &MeilisearchBookRepository{}
}
