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

func (service *newsApiService) GetEverything(query string, page int, pageSize int) (*newsapi.Everything, error) {
	return service.repository.GetEverything(query, page, pageSize)
}
