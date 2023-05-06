package usecases

import (
	"news-api/internal/domain/entities/newsapi"
	"news-api/internal/domain/services"
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

func (usecase *NewsApiUsecase) GetEverything(query string, page int, pageSize int) (*newsapi.Everything, error) {
	return usecase.service.GetEverything(query, page, pageSize)
}
