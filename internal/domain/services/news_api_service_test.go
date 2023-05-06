package services

import (
	"news-api/internal/domain/entities/newsapi"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) GetEverything(query string, page int, pageSize int) (*newsapi.Everything, error) {
	args := m.Called(query, page, pageSize)
	return args.Get(0).(*newsapi.Everything), args.Error(1)
}

// 正常終了ケース
func TestGetEverything(t *testing.T) {
	mockRepo := &mockRepository{}
	service := NewNewsApiService(mockRepo)

	query := "test"
	page := 1
	pageSize := 10

	expectedEverything := &newsapi.Everything{}
	mockRepo.On("GetEverything", query, page, pageSize).Return(expectedEverything, nil).Once()

	_, err := service.GetEverything("test", 1, 10)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
