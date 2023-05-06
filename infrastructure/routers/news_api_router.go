package routers

import (
	"news-api/domain/services"
	"news-api/domain/usecases"
	"news-api/infrastructure"
	"news-api/infrastructure/apiclient"
	"news-api/infrastructure/persistence"
	"news-api/interfaces/controller"
	"news-api/utils"

	"github.com/gin-gonic/gin"
)

// NewAPIのURLドメイン
const newsAPIBaseURL = "https://newsapi.org"

// @summary		NewsAPIの記事を取得する
// @description	NewsAPIの記事を取得する( https://newsapi.org/docs/endpoints/everything のエンドポイントを使用します)
// @tags		news-api
// @produce		json
// @param		query query string true "query"
// @param		page query int true "page"
// @param		pageSize query int true "pageSize"
// @success 200 {object} controller.NewsAPIResonse "NewsAPIの記事情報を返します"
// @failure 400 {object} controller.ErrorResponse "バリデーションエラーメッセージを返します"
// @failure 500 {object} controller.ErrorResponse "システムエラーメッセージを返します"
// @router		/news-api [get]
func setNewsAPIRouter(router *gin.Engine) {
	router.GET("/news-api", func(c *gin.Context) {
		httpStatusCode, data := getController().GetEverything(c)
		c.JSON(httpStatusCode, gin.H(utils.StructToMap(data)))
	})
}

func getController() *controller.NewsAPIController {
	externalAPIClient := apiclient.NewExternalClient(
		newsAPIBaseURL,
		infrastructure.NewSentryErrorHandler(),
		apiclient.NewHTTPClient(),
	)
	newsAPIRepository := persistence.NewNewsAPIRepositoryImpl(externalAPIClient)
	newsAPIService := services.NewNewsApiService(newsAPIRepository)
	newsAPIUsecase := usecases.NewNewsApiUsecase(newsAPIService)
	return controller.NewNewsAPIController(newsAPIUsecase)
}
