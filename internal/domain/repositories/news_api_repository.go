package repositories

import "news-api/internal/domain/entities/newsapi"

type NewsAPIRepositoryInterface interface {
	GetEverything(query string, page int, pageSize int) (*newsapi.Everything, error)
}
