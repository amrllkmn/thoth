package search

import "github.com/amrllkmn/thoth/backend/internal/utils"

type SQLiteSearchService struct {
	repo SQLiteBookRepository
}

func (s *SQLiteSearchService) FindAll()                 {}
func (s *SQLiteSearchService) FindByQuery(query string) {}
func (s *SQLiteSearchService) FindByID(id uint)         {}

func NewSQLiteSearchService(repo SQLiteBookRepository) utils.SearchService {
	return &SQLiteSearchService{repo: repo}
}
