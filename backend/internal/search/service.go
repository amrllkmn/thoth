package search

import "github.com/amrllkmn/thoth/backend/internal/utils"

type SQLiteSearchService struct {
	repo utils.BookRepository
}

func (s *SQLiteSearchService) FindAll(page, limit int) ([]utils.Book, error) {
	books, err := s.repo.FindAll(page, limit)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *SQLiteSearchService) FindByQuery(query string, page, limit int) ([]utils.Book, error) {
	books, err := s.repo.FindByQuery(query, page, limit)
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (s *SQLiteSearchService) FindByID(isbn string) (*utils.Book, error) {
	book, err := s.repo.FindByID(isbn)
	return book, err
}

func NewSQLiteSearchService(repo utils.BookRepository) utils.SearchService {
	return &SQLiteSearchService{repo: repo}
}
