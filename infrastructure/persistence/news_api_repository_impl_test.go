package persistence

import (
	"errors"
	"net/http"
	"news-api/domain/entities/newsapi"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockExternalClient struct {
	mock.Mock
}

func (m *mockExternalClient) Get(
	path string,
	headers http.Header,
	queryParams map[string]string,
	response interface{},
) error {
	args := m.Called(path, headers, queryParams, response)
	return args.Error(0)
}

// 正常終了ケース
func TestGetEverything(t *testing.T) {
	newsApiKey := "test"
	page := 1
	pageSize := 20

	os.Setenv("NEWS_API_KEY", newsApiKey)

	mockExternalClient := &mockExternalClient{}
	mockClientOn(
		mockExternalClient,
		newsApiKey,
		page,
		pageSize,
	).Return(nil)

	repositoryImpl := NewNewsAPIRepositoryImpl(mockExternalClient)
	_, err := repositoryImpl.GetEverything(newsApiKey, page, pageSize)

	// エラーが発生していないこと
	assert.Nil(t, err)
	// Getメソッドが呼ばれていること
	mockExternalClient.AssertExpectations(t)
}

// 異常終了ケース
func TestGetEverythingError(t *testing.T) {
	newsApiKey := "test"
	page := 1
	pageSize := 20

	os.Setenv("NEWS_API_KEY", newsApiKey)

	mockExternalClient := &mockExternalClient{}
	mockClientOn(
		mockExternalClient,
		newsApiKey,
		page,
		pageSize,
	).Return(errors.New("test"))

	repositoryImpl := NewNewsAPIRepositoryImpl(mockExternalClient)
	_, err := repositoryImpl.GetEverything(newsApiKey, page, pageSize)

	// エラーが発生していること
	assert.NotNil(t, err)
	// Getメソッドが呼ばれていること
	mockExternalClient.AssertExpectations(t)
}

// Getメソッドのモックを設定
func mockClientOn(
	mockExternalClient *mockExternalClient,
	newsApiKey string,
	page int,
	pageSize int,
) *mock.Call {
	return mockExternalClient.On(
		"Get",
		"/v2/everything",
		http.Header{
			"X-Api-Key": {"test"},
		},
		map[string]string{
			"q":        newsApiKey,
			"page":     strconv.Itoa(page),
			"language": "jp",
			"pageSize": strconv.Itoa(pageSize),
		},
		&newsapi.Everything{},
	)
}
