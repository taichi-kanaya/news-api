package apiclient

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"news-api/internal/config"
	"news-api/internal/domain/entities/newsapi"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// HttpClientのモック
type mockHttpClient struct {
	mock.Mock
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

// Sentryクライアントのモック
type mockSentry struct {
	mock.Mock
}

func (m *mockSentry) SendSentry(err error) {
	m.Called(err)
	return
}

// io.ReadAllメソッドで擬似的にエラーを発生させるための構造体
type errorReader struct{}

func (e *errorReader) Read(p []byte) (int, error) {
	return 0, errors.New("custom error from Read")
}

// Getリクエストのパラメータを生成
func createGetRequestParams() (string, map[string][]string, map[string]string, *newsapi.Everything) {
	path := "/test"
	headers := map[string][]string{
		"X-Api-Key": {"abcde"},
	}
	queryParams := map[string]string{
		"q":        "test",
		"page":     "1",
		"language": "jp",
		"pageSize": "20",
	}
	everything := &newsapi.Everything{}
	return path, headers, queryParams, everything
}

// 正常終了ケース
func TestGet(t *testing.T) {
	// HttpClientのモックを設定
	mockHttpClient := &mockHttpClient{}
	mockHttpClient.DoFunc = func(req *http.Request) (*http.Response, error) {
		// Getリクエスト結果の期待値を設定
		res := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"status":"ok"}`)),
		}
		return res, nil
	}

	// 各種パラメータ設定
	client := NewExternalClient(
		"http://localhost:8080",
		&mockSentry{},
		mockHttpClient,
	)
	path, headers, queryParams, _ := createGetRequestParams()
	everything := &newsapi.Everything{}

	// テスト対象メソッド実行
	err := client.Get(path, headers, queryParams, everything)

	// エラーが発生していないこと
	assert.Nil(t, err)
	// Getリクエスト結果が正しいこと
	assert.Equal(t, "ok", everything.Status)
}

// HTTPリクエスト生成に失敗した場合
func TestGetHttpRequestErr(t *testing.T) {
	mockHttpClient := &mockHttpClient{}
	mockHttpClient.DoFunc = func(req *http.Request) (*http.Response, error) {
		return nil, nil
	}

	// Sentryクライアントのモックを設定
	mockSentry := &mockSentry{}
	mockSentry.On("SendSentry", mock.Anything).Return()

	client := NewExternalClient(
		"", // 制御文字（= DEL）を入れてエラーを発生させる
		mockSentry,
		mockHttpClient,
	)

	err := client.Get(createGetRequestParams())

	// エラーが発生していること
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"%d:%s",
		http.StatusInternalServerError,
		config.ERROR_MESSAGE_500,
	), err.Error())
	// Sentryにエラーが送信されていること
	mockSentry.AssertExpectations(t)
}

// 外部APIコールに失敗した場合
func TestGetDoErr(t *testing.T) {
	mockHttpClient := &mockHttpClient{}
	mockHttpClient.DoFunc = func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("error") // 外部APIコールでエラー発生したと仮定
	}

	mockSentry := &mockSentry{}
	mockSentry.On("SendSentry", mock.Anything).Return()

	client := NewExternalClient(
		"http://localhost:8080",
		mockSentry,
		mockHttpClient,
	)

	err := client.Get(createGetRequestParams())

	// エラーが発生していること
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"%d:%s",
		http.StatusInternalServerError,
		config.ERROR_MESSAGE_500,
	), err.Error())
	// Sentryにエラーが送信されていること
	mockSentry.AssertExpectations(t)
}

// レスポンス読み込みに失敗した場合
func TestReadResponseErr(t *testing.T) {
	mockHttpClient := &mockHttpClient{}
	mockHttpClient.DoFunc = func(req *http.Request) (*http.Response, error) {
		res := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(&errorReader{}), // レスポンス読み込みでエラー発生したと仮定
		}
		return res, nil
	}

	mockSentry := &mockSentry{}
	mockSentry.On("SendSentry", mock.Anything).Return()

	client := NewExternalClient(
		"http://localhost:8080",
		mockSentry,
		mockHttpClient,
	)
	err := client.Get(createGetRequestParams())

	// エラーが発生していること
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"%d:%s",
		http.StatusInternalServerError,
		config.ERROR_MESSAGE_500,
	), err.Error())
	// Sentryにエラーが送信されていること
	mockSentry.AssertExpectations(t)
}

// レスポンスのステータスコードが200以外の場合
func TestStatusCodeErr(t *testing.T) {
	mockHttpClient := &mockHttpClient{}
	mockHttpClient.DoFunc = func(req *http.Request) (*http.Response, error) {
		res := &http.Response{
			StatusCode: http.StatusServiceUnavailable,
			Body:       io.NopCloser(strings.NewReader(`{"status":"ng","message":"503 error"}`)),
		}
		return res, nil
	}

	mockSentry := &mockSentry{}
	mockSentry.On("SendSentry", mock.Anything).Return()

	client := NewExternalClient(
		"http://localhost:8080",
		mockSentry,
		mockHttpClient,
	)
	err := client.Get(createGetRequestParams())

	// エラーが発生していること
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"%d:%s",
		http.StatusInternalServerError,
		config.ERROR_MESSAGE_500,
	), err.Error())
	// Sentryにエラーが送信されていること
	mockSentry.AssertExpectations(t)
}

// レスポンスのJSONのパースに失敗した場合

func TestJsonParseErr(t *testing.T) {
	mockHttpClient := &mockHttpClient{}
	mockHttpClient.DoFunc = func(req *http.Request) (*http.Response, error) {
		res := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"status"::"ok"}`)),
		}
		return res, nil
	}

	mockSentry := &mockSentry{}
	mockSentry.On("SendSentry", mock.Anything).Return()

	client := NewExternalClient(
		"http://localhost:8080",
		mockSentry,
		mockHttpClient,
	)
	err := client.Get(createGetRequestParams())

	// エラーが発生していること
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"%d:%s",
		http.StatusInternalServerError,
		config.ERROR_MESSAGE_500,
	), err.Error())
	// Sentryにエラーが送信されていること
	mockSentry.AssertExpectations(t)
}
