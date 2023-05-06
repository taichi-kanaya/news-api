package controller

import (
	"net/http"
	"net/http/httptest"
	"news-api/config"
	"news-api/domain"
	"news-api/domain/entities/newsapi"
	"news-api/validation"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUsecase struct {
	mock.Mock
}

func (m *mockUsecase) GetEverything(
	query string,
	page int,
	pageSize int,
) (*newsapi.Everything, error) {
	args := m.Called(query, page, pageSize)
	return args.Get(0).(*newsapi.Everything), args.Error(1)
}

// 正常終了ケース
func TestGetEverything(t *testing.T) {
	// NewsAPIから取得した記事として扱うテストデータ
	testData := &newsapi.Everything{
		Status:       "ok",
		TotalResults: 2,
		Articles: []newsapi.Article{
			{
				Source: newsapi.ArticleSource{
					Id:   "source-1",
					Name: "Source One",
				},
				Author:      "Author 1",
				Title:       "Title 1",
				Description: "Description 1",
				URL:         "http://localhost:8080/article1",
				URLToImage:  "http://localhost:8080/image1.jpg",
				PublishedAt: "2023-05-06T12:00:00Z",
				Content:     "Content 1",
			},
			{
				Source: newsapi.ArticleSource{
					Id:   "source-2",
					Name: "Source Two",
				},
				Author:      "Author 2",
				Title:       "Title 2",
				Description: "Description 2",
				URL:         "http://localhost:8080/article2",
				URLToImage:  "http://localhost:8080/image2.jpg",
				PublishedAt: "2023-05-06T13:00:00Z",
				Content:     "Content 2",
			},
		},
	}

	// usecaseのモックを設定
	mockUsecase := &mockUsecase{}
	mockUsecase.On(
		"GetEverything",
		"test",
		1,
		10,
	).Return(
		testData,
		nil,
	).Once()

	newsApiController := NewNewsAPIController(mockUsecase)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	validation.RegisterValidation()
	c.Request = httptest.NewRequest("GET", "/news-api?query=test&page=1&pageSize=10", nil)

	// テスト対象コードを実行
	httpStatusCode, newsAPIResponse := newsApiController.GetEverything(c)

	// エラーが発生していないこと
	assert.Equal(t, httpStatusCode, http.StatusOK)
	// テストデータと返却値が合致していること
	assert.Equal(t, newsAPIResponse, &NewsAPIResonse{
		TotalResults: testData.TotalResults,
		Articles:     testData.Articles,
	})
	// GetEverythingメソッドが呼び出されていること
	mockUsecase.AssertExpectations(t)
}

// usecaseがエラーを返してきた場合
func TestUsecaseError(t *testing.T) {
	mockUsecase := &mockUsecase{}
	mockUsecase.On(
		"GetEverything",
		"test",
		1,
		10,
	).Return(
		&newsapi.Everything{},
		domain.NewCustomError(
			http.StatusInternalServerError,
			[]string{config.ERROR_MESSAGE_500},
		),
	).Once()

	newsApiController := NewNewsAPIController(mockUsecase)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	validation.RegisterValidation()
	c.Request = httptest.NewRequest("GET", "/news-api?query=test&page=1&pageSize=10", nil)

	httpStatusCode, err := newsApiController.GetEverything(c)

	// エラーが発生していること
	assert.Equal(
		t,
		httpStatusCode,
		http.StatusInternalServerError,
	)
	assert.Equal(
		t,
		err.(*ErrorResponse).Errors,
		[]string{config.ERROR_MESSAGE_500},
	)
	// GetEverythingメソッドが呼び出されていること
	mockUsecase.AssertExpectations(t)
}

// 各種パラメータエラー
func TestParamsError(t *testing.T) {
	mockUsecase := &mockUsecase{}
	newsApiController := NewNewsAPIController(mockUsecase)

	testCases := []struct {
		name           string
		requestURL     string
		expectedStatus int
		expectedError  string
	}{
		{
			name:          "queryが未指定",
			requestURL:    "/news-api?page=1&pageSize=10",
			expectedError: "queryを指定してください",
		},
		{
			name:          "queryが500文字超え",
			requestURL:    "/news-api?query=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxy&page=1&pageSize=10",
			expectedError: "queryには500文字以内で指定してください",
		},
		{
			name:          "pageが未指定",
			requestURL:    "/news-api?query=test&pageSize=10",
			expectedError: "pageを指定してください",
		},
		{
			name:          "pageが文字列",
			requestURL:    "/news-api?query=test&page=a&pageSize=10",
			expectedError: "pageには1以上の数値を指定してください",
		},
		{
			name:          "pageが1未満",
			requestURL:    "/news-api?query=test&page=0&pageSize=10",
			expectedError: "pageには1以上の数値を指定してください",
		},
		{
			name:          "pageSizeが未指定",
			requestURL:    "/news-api?query=test&page=1",
			expectedError: "pageSizeを指定してください",
		},
		{
			name:          "pageSizeが文字列",
			requestURL:    "/news-api?query=test&page=1&pageSize=a",
			expectedError: "pageSizeには1以上の数値を指定してください",
		},
		{
			name:          "pageSizeが1未満",
			requestURL:    "/news-api?query=test&page=1&pageSize=0",
			expectedError: "pageSizeには1以上の数値を指定してください",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			validation.RegisterValidation()
			c.Request = httptest.NewRequest("GET", tc.requestURL, nil)

			httpStatusCode, err := newsApiController.GetEverything(c)

			// エラーが発生していること
			assert.Equal(t, http.StatusBadRequest, httpStatusCode)
			assert.Equal(t, tc.expectedError, err.(*ErrorResponse).Errors[0])
			// GetEverythingメソッドが呼び出されていないこと
			mockUsecase.AssertNotCalled(t, "GetEverything")
		})
	}
}
