package indexed_search

import "github.com/amrllkmn/thoth/backend/internal/utils"

type MeilisearchSearchService struct {
	repo utils.BookRepository
}

func (s *MeilisearchSearchService) FindAll(page, limit int) ([]utils.Book, error) {
	books, err := s.repo.FindAll(page, limit)
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (s *MeilisearchSearchService) FindByQuery(query string, page, limit int) ([]utils.Book, error) {
	books, err := s.repo.FindByQuery(query, page, limit)
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (s *MeilisearchSearchService) FindByID(isbn string) (*utils.Book, error) {
	book, err := s.repo.FindByID(isbn)
	return book, err
}

func NewMeilisearchSearchService(repo utils.BookRepository) utils.SearchService {
	return &MeilisearchSearchService{repo: repo}
}
