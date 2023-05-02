package usecases

import (
	"news-api/internal/domain/entities/news_api"
	"news-api/internal/domain/services"
)

type NewsApiUsecase struct {
	service *services.NewsApiService
}

func NewNewsApiUsecase(service *services.NewsApiService) *NewsApiUsecase {
	return &NewsApiUsecase{service: service}
}

func (usecase *NewsApiUsecase) GetEverything(query string, page int) (*news_api.Everything, error) {
	return usecase.service.GetEverything(query, page)
}
