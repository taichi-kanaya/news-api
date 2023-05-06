package persistence

import (
	"news-api/internal/config"
	"news-api/internal/domain/entities/newsapi"
	"news-api/internal/domain/repositories"
	"news-api/internal/infrastructure/apiclient"
	"strconv"
)

type NewsAPIRepositoryImpl struct {
	externalApi apiclient.ExternalAPIInterface
}

func NewNewsAPIRepositoryImpl(
	externalApi apiclient.ExternalAPIInterface,
) repositories.NewsAPIRepositoryInterface {
	return &NewsAPIRepositoryImpl{externalApi: externalApi}
}

// NewsAPIの記事を検索する
// ref: https://newsapi.org/docs/endpoints/everything
func (n *NewsAPIRepositoryImpl) GetEverything(
	query string,
	page int,
	pageSize int,
) (*newsapi.Everything, error) {
	headers := map[string][]string{
		"X-Api-Key": {config.GetNewsAPIKey()},
	}
	queryParams := map[string]string{
		"q":        query,
		"page":     strconv.Itoa(page),
		"language": "jp",
		"pageSize": strconv.Itoa(pageSize),
	}
	everything := &newsapi.Everything{}
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
