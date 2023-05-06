package repositories

import "news-api/domain/entities/newsapi"

// NewsAPIの記事を検索する
type NewsAPIRepositoryInterface interface {
	GetEverything(query string, page int, pageSize int) (*newsapi.Everything, error)
}
