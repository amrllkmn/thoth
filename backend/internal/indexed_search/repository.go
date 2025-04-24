package indexed_search

import (
	"encoding/json"
	"fmt"

	"github.com/amrllkmn/thoth/backend/internal/utils"
	"github.com/meilisearch/meilisearch-go"

	"github.com/labstack/gommon/log"
)

type MeilisearchBookRepository struct {
	ms_client meilisearch.ServiceManager
}

func (r *MeilisearchBookRepository) FindAll(page, limit int) ([]utils.Book, error) {
	var books []utils.Book
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 20
	}

	offset := (page - 1) * limit
	res, err := r.ms_client.Index("books").Search("", &meilisearch.SearchRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Error("Searching failed...")
		return nil, err
	}

	hitsJSON, err := json.Marshal(res.Hits)
	if err != nil {
		log.Error("Marshalling failed...")
		return nil, err
	}

	err = json.Unmarshal(hitsJSON, &books)
	if err != nil {
		log.Error("Unmarshalling failed...")
	}

	return books, nil

}
func (r *MeilisearchBookRepository) FindByQuery(query string, page, limit int) ([]utils.Book, error) {
	var books []utils.Book
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 20
	}

	offset := (page - 1) * limit
	res, err := r.ms_client.Index("books").Search(query, &meilisearch.SearchRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Error("Searching failed...")
		return nil, err
	}

	hitsJSON, err := json.Marshal(res.Hits)
	if err != nil {
		log.Error("Marshalling failed...")
		return nil, err
	}

	err = json.Unmarshal(hitsJSON, &books)
	if err != nil {
		log.Error("Unmarshalling failed...")
	}

	return books, nil
}
func (r *MeilisearchBookRepository) FindByID(isbn string) (*utils.Book, error) {
	var books []utils.Book
	filter := fmt.Sprintf(`isbn13 = "%s" OR isbn10 = "%s"`, isbn, isbn)
	res, err := r.ms_client.Index("books").Search("", &meilisearch.SearchRequest{
		Filter: filter,
		Limit:  1,
	})

	if err != nil {
		log.Error("Searching failed...")
		return nil, err
	}

	hitsJSON, err := json.Marshal(res.Hits)
	if err != nil {
		log.Error("Marshalling failed...")
		return nil, err
	}

	err = json.Unmarshal(hitsJSON, &books)
	if err != nil {
		log.Error("Unmarshalling failed...")
	}

	if len(books) == 0 {
		return nil, nil
	}

	return &books[0], nil
}

func NewMeilisearchBookRepository(ms meilisearch.ServiceManager) utils.BookRepository {
	return &MeilisearchBookRepository{
		ms_client: ms,
	}
}
