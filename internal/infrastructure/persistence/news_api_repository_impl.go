package persistence

import (
	"news-api/internal/config"
	"news-api/internal/domain/entities/news_api"
	"news-api/internal/domain/repositories"
	"news-api/internal/infrastructure/api_client"
	"strconv"
)

type NewsAPIRepositoryImpl struct {
	externalApi *api_client.ExternalClient
}

func NewNewsAPIRepositoryImpl(externalApi *api_client.ExternalClient) repositories.NewsAPIRepository {
	return &NewsAPIRepositoryImpl{externalApi: externalApi}
}

// NewsAPIの記事を検索する
// ref: https://newsapi.org/docs/endpoints/everything
func (n *NewsAPIRepositoryImpl) GetEverything(
	query string,
	page int,
	pageSize int,
) (*news_api.Everything, error) {
	headers := map[string][]string{
		"X-Api-Key": {config.GetNewsAPIKey()},
	}
	queryParams := map[string]string{
		"q":        query,
		"page":     strconv.Itoa(page),
		"language": "jp",
		"pageSize": strconv.Itoa(pageSize),
	}
	everything := &news_api.Everything{}
	err := n.externalApi.Get(
		"/v2/everything",
		headers,
		queryParams,
		everything,
	)
	if err != nil {
		return nil, err
	}
	return everything, nil
}
