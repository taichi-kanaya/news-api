package repositories

import "news-api/internal/domain/entities/news_api"

type NewsAPIRepository interface {
	GetEverything(query string, page int) (*news_api.Everything, error)
}
