package usecases

import (
	"news-api/internal/domain/entities/newsapi"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func (m *mockService) GetEverything(query string, page int, pageSize int) (*newsapi.Everything, error) {
	args := m.Called(query, page, pageSize)
	return args.Get(0).(*newsapi.Everything), args.Error(1)
}

// 正常終了ケース
func TestNewNewsApiUsecase(t *testing.T) {
	mockService := &mockService{}
	usecase := NewNewsApiUsecase(mockService)

	query := "test"
	page := 1
	pageSize := 10

	expectedEverything := &newsapi.Everything{}
	mockService.On("GetEverything", query, page, pageSize).Return(expectedEverything, nil).Once()

	_, err := usecase.GetEverything(query, page, pageSize)

	assert.Nil(t, err)
	mockService.AssertExpectations(t)
}
