/*
NewsAPIから記事を取得してユーザに返すコントローラ
*/
package controller

import (
	"fmt"
	"net/http"
	"news-api/domain"
	"news-api/domain/entities/newsapi"
	"news-api/domain/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NewsAPIController struct {
	usecase usecases.NewsApiUsecaseInterface
}

func NewNewsAPIController(usecase usecases.NewsApiUsecaseInterface) *NewsAPIController {
	return &NewsAPIController{usecase: usecase}
}

// ユーザに返す正常時レスポンス
type NewsAPIResonse struct {
	TotalResults int               `json:"totalResults"`
	Articles     []newsapi.Article `json:"articles"`
}

// HTTPリクエストパラメータ
type RequestParam struct {
	Query    string `form:"query" binding:"required,max=500"`
	Page     string `form:"page" binding:"required,string-min-value=1"`
	PageSize string `form:"pageSize" binding:"required,string-min-value=1"`
}

// NewsAPIから記事を取得してユーザに返す
//
// Parameters:
//   - c: *gin.Context
//
// Returns:
//   - int: HTTPステータスコード
//   - interface{}: レスポンスボディ(NewsAPIResonse or ErrorResponse)
func (newsAPIController *NewsAPIController) GetEverything(c *gin.Context) (int, interface{}) {
	// リクエストパラメータをバインド
	var requestParam RequestParam
	if err := bind(c, &requestParam); err != nil {
		return err.(*domain.CustomError).HttpStatusCode, &ErrorResponse{
			Errors: err.(*domain.CustomError).Messages,
		}
	}

	// リクエストパラメータを元にNewsAPIから記事を取得
	page, _ := strconv.Atoi(requestParam.Page)
	pageSize, _ := strconv.Atoi(requestParam.PageSize)
	everything, err := newsAPIController.usecase.GetEverything(requestParam.Query, page, pageSize)
	if err != nil {
		return err.(*domain.CustomError).HttpStatusCode, &ErrorResponse{
			Errors: err.(*domain.CustomError).Messages,
		}
	}

	return http.StatusOK, &NewsAPIResonse{
		TotalResults: everything.TotalResults,
		Articles:     everything.Articles,
	}
}

// リクエストパラメータをバインドする
//
// Parameters:
//   - c: *gin.Context
//   - requestParam: *RequestParam
//
// Returns:
//   - error: エラー情報
func bind(c *gin.Context, requestParam *RequestParam) error {
	if err := c.ShouldBind(&requestParam); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, e := range validationErrors {
			fieldName := e.Field()
			tag := e.Tag()

			switch fieldName {
			case "Query":
				if tag == "required" {
					errorMessages[i] = fmt.Sprintf("queryを指定してください")
				} else if tag == "max" {
					errorMessages[i] = fmt.Sprintf("queryには500文字以内で指定してください")
				}
			case "Page":
				if tag == "required" {
					errorMessages[i] = fmt.Sprintf("pageを指定してください")
				} else if tag == "string-min-value" {
					errorMessages[i] = fmt.Sprintf("pageには1以上の数値を指定してください")
				}
			case "PageSize":
				if tag == "required" {
					errorMessages[i] = fmt.Sprintf("pageSizeを指定してください")
				} else if tag == "string-min-value" {
					errorMessages[i] = fmt.Sprintf("pageSizeには1以上の数値を指定してください")
				}
			}
		}
		if len(errorMessages) > 0 {
			return domain.NewCustomError(
				http.StatusBadRequest,
				errorMessages,
			)
		}
	}
	return nil
}
