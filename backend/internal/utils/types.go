package utils

type Book struct {
	ID            uint    `gorm:"primaryKey"`
	Isbn13        string  `gorm:"unique" json:"isbn13"`
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
	FindAll(page, limit int) ([]Book, error)
	FindByQuery(query string, page, limit int) ([]Book, error)
	FindByID(isbn string) (*Book, error)
}

type SearchService interface {
	FindAll(page, limit int) ([]Book, error)
	FindByQuery(query string, page, limit int) ([]Book, error)
	FindByID(isbn string) (*Book, error)
}
