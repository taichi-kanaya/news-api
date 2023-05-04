package services

import (
	"news-api/internal/domain/entities/news_api"
	"news-api/internal/domain/repositories"
)

type NewsApiService struct {
	repository repositories.NewsAPIRepository
}

func NewNewsApiService(repository repositories.NewsAPIRepository) *NewsApiService {
	return &NewsApiService{repository: repository}
}

func (service *NewsApiService) GetEverything(query string, page int, pageSize int) (*news_api.Everything, error) {
	return service.repository.GetEverything(query, page, pageSize)
}
