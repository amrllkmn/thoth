package search

type SearchService interface {
	FindAll()
	FindByQuery(query string)
	FindByID(id uint)
}

type SQLiteSearchService struct {
	repo SQLiteBookRepository
}

func (s *SQLiteSearchService) FindAll()                 {}
func (s *SQLiteSearchService) FindByQuery(query string) {}
func (s *SQLiteSearchService) FindByID(id uint)         {}

func NewSQLiteSearchService(repo SQLiteBookRepository) SearchService {
	return &SQLiteSearchService{repo: repo}
}

type MeilisearchSearchService struct {
	repo MeilisearchBookRepository
}

func (s *MeilisearchSearchService) FindAll()                 {}
func (s *MeilisearchSearchService) FindByQuery(query string) {}
func (s *MeilisearchSearchService) FindByID(id uint)         {}

func NewMeilisearchSearchService(repo MeilisearchBookRepository) SearchService {
	return &MeilisearchSearchService{repo: repo}
}
