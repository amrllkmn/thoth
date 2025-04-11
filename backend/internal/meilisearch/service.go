package meilisearch

type MeilisearchSearchService struct {
	repo MeilisearchBookRepository
}

func (s *MeilisearchSearchService) FindAll()                 {}
func (s *MeilisearchSearchService) FindByQuery(query string) {}
func (s *MeilisearchSearchService) FindByID(id uint)         {}

// func NewMeilisearchSearchService(repo MeilisearchBookRepository) utils.SearchService {
// 	return &MeilisearchSearchService{repo: repo}
// }
