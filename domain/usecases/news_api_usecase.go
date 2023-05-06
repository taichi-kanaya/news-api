package usecases

import (
	"news-api/domain/entities/newsapi"
	"news-api/domain/services"
)

type NewsApiUsecaseInterface interface {
	GetEverything(query string, page int, pageSize int) (*newsapi.Everything, error)
}

type NewsApiUsecase struct {
	service services.NewsApiServiceInterface
}

func NewNewsApiUsecase(service services.NewsApiServiceInterface) NewsApiUsecaseInterface {
	return &NewsApiUsecase{service: service}
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
func (usecase *NewsApiUsecase) GetEverything(query string, page int, pageSize int) (*newsapi.Everything, error) {
	return usecase.service.GetEverything(query, page, pageSize)
}
