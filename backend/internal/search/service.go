package search

import "github.com/amrllkmn/thoth/backend/internal/utils"

type SQLiteSearchService struct {
	repo utils.BookRepository
}

func (s *SQLiteSearchService) FindAll() ([]utils.Book, error) {
	books, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *SQLiteSearchService) FindByQuery(query string) {}
func (s *SQLiteSearchService) FindByID(id uint)         {}

func NewSQLiteSearchService(repo utils.BookRepository) utils.SearchService {
	return &SQLiteSearchService{repo: repo}
}
