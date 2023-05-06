package services

import (
	"news-api/domain/entities/newsapi"
	"news-api/domain/repositories"
)

type NewsApiServiceInterface interface {
	GetEverything(query string, page int, pageSize int) (*newsapi.Everything, error)
}

type newsApiService struct {
	repository repositories.NewsAPIRepositoryInterface
}

func NewNewsApiService(repository repositories.NewsAPIRepositoryInterface) NewsApiServiceInterface {
	return &newsApiService{repository: repository}
}

// NewsAPIの記事を検索する
//
// Parameters:
//   - query: 検索クエリ
//   - page: ページ番号
//   - pageSize: 1ページあたりの記事数
//
// Returns:
//   - *newsapi.Everything
//   - error
func (service *newsApiService) GetEverything(query string, page int, pageSize int) (*newsapi.Everything, error) {
	return service.repository.GetEverything(query, page, pageSize)
}
