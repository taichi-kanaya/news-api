package controller

import (
	"net/http"
	"news-api/internal/domain"
	"news-api/internal/domain/entities/news_api"
	"news-api/internal/domain/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NewsAPIController struct {
	usecase *usecases.NewsApiUsecase
}

func NewNewsAPIController(usecase *usecases.NewsApiUsecase) *NewsAPIController {
	return &NewsAPIController{usecase: usecase}
}

type NewsAPIResonse struct {
	TotalResults int                `json:"totalResults"`
	Articles     []news_api.Article `json:"articles"`
}

func (controller *NewsAPIController) GetEverything(c *gin.Context) (*NewsAPIResonse, error) {
	// 検索ワードを取得
	query := c.Query("query")

	// ページ数を取得
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	// バリデーションチェック
	if len(query) <= 0 {
		return nil, domain.NewCustomError(
			http.StatusBadRequest,
			"queryパラメータを指定してください",
		)
	} else if len(query) > 500 {
		return nil, domain.NewCustomError(
			http.StatusBadRequest,
			"queryパラメータは500文字以内で指定してください",
		)
	}

	everything, err := controller.usecase.GetEverything(query, page)
	if err != nil {
		return nil, err
	}
	response := &NewsAPIResonse{
		TotalResults: everything.TotalResults,
		Articles:     everything.Articles,
	}
	return response, nil
}
