package meilisearch

type MeilisearchBookRepository struct {
}

func (r *MeilisearchBookRepository) FindAll()                 {}
func (r *MeilisearchBookRepository) FindByQuery(query string) {}
func (r *MeilisearchBookRepository) FindByID(id uint)         {}

// func NewMeilisearchBookRepository() utils.BookRepository {
// 	return &MeilisearchBookRepository{}
// }
