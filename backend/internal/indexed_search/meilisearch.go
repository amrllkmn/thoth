package indexed_search

import (
	"os"

	"github.com/meilisearch/meilisearch-go"
)

func InitMeilisearchClient() meilisearch.ServiceManager {
	url := os.Getenv("MEILISEARCH_URL")
	api_key := os.Getenv("MEILISEARCH_API_KEY")
	if url == "" {
		url = "http://localhost:7700" // Default Meilisearch URL
	}

	if api_key == "" {
		panic("MEILISEARCH_API_KEY is not set")
	}

	return meilisearch.New(url, meilisearch.WithAPIKey(api_key))

}
