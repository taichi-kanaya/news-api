package services

import (
	"news-api/internal/domain/entities/newsapi"
	"news-api/internal/domain/repositories"
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

func (service *newsApiService) GetEverything(query string, page int, pageSize int) (*newsapi.Everything, error) {
	return service.repository.GetEverything(query, page, pageSize)
}
