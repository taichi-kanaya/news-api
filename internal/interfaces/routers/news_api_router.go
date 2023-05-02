package routers

import (
	"net/http"
	"news-api/internal/domain"
	"news-api/internal/domain/services"
	"news-api/internal/domain/usecases"
	"news-api/internal/infrastructure/api_client"
	"news-api/internal/infrastructure/persistence"
	"news-api/internal/interfaces/controller"

	"github.com/gin-gonic/gin"
)

// NewAPIのURLドメイン
const newsAPIBaseURL = "https://newsapi.org"

// @summary		NewsAPIの記事を取得する
// @description	NewsAPIの記事を取得する( https://newsapi.org/docs/endpoints/everything のエンドポイントを使用します)
// @tags		news-api
// @produce		json
// @param		query query string true "query"
// @param		page query int false "page"
// @success 200 {object} string "NewsAPIの記事情報を返します"
// @failure 400 {object} string "バリデーションエラーメッセージを返します"
// @router		/news-api [get]
func setNewsAPIRouter(router *gin.Engine) {
	router.GET("/news-api", func(c *gin.Context) {
		res, err := getController().GetEverything(c)
		httpStatusCode := http.StatusOK
		data := map[string]any{
			"data": res,
		}
		if err != nil {
			httpStatusCode = err.(*domain.CustomError).HttpStatusCode
			data = map[string]any{
				"error": err.(*domain.CustomError).Message,
			}
		}
		c.JSON(httpStatusCode, gin.H(data))
	})
}

func getController() *controller.NewsAPIController {
	externalAPIClient := api_client.NewExternalClient(newsAPIBaseURL)
	newsAPIRepository := persistence.NewNewsAPIRepositoryImpl(externalAPIClient)
	newsAPIService := services.NewNewsApiService(newsAPIRepository)
	newsAPIUsecase := usecases.NewNewsApiUsecase(newsAPIService)
	return controller.NewNewsAPIController(newsAPIUsecase)
}
