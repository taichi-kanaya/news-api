package api_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"news-api/internal/config"
	"news-api/internal/domain"
	"news-api/internal/infrastructure"
)

type ExternalClient struct {
	BaseURL string
}

func NewExternalClient(baseURL string) *ExternalClient {
	return &ExternalClient{BaseURL: baseURL}
}

// 外部APIのGETリクエストを実行する
func (c *ExternalClient) Get(
	path string,
	headers http.Header,
	queryParams map[string]string,
	response interface{},
) error {

	// URLクエリストリング生成
	url := c.BaseURL + path
	if queryParams != nil {
		url += "?"
		for k, v := range queryParams {
			url += fmt.Sprintf("%s=%s&", k, v)
		}
		url = url[:len(url)-1]
	}

	// HTTPリクエスト生成
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		infrastructure.NewSentryErrorHandler().SendSentry(
			fmt.Errorf("外部APIのGETリクエスト生成に失敗しました。url: %s", url),
		)
		return domain.NewCustomError(
			http.StatusInternalServerError,
			[]string{config.ERROR_MESSAGE_500},
		)
	}

	// HTTPヘッダ設定
	if headers != nil {
		req.Header = headers
	}

	// 外部APIコール
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		infrastructure.NewSentryErrorHandler().SendSentry(
			fmt.Errorf("外部APIのGETリクエスト実行に失敗しました。url: %s", url),
		)
		return domain.NewCustomError(
			http.StatusInternalServerError,
			[]string{config.ERROR_MESSAGE_500},
		)
	}
	defer res.Body.Close()

	// レスポンス読み込み
	body, err := io.ReadAll(res.Body)
	if err != nil {
		infrastructure.NewSentryErrorHandler().SendSentry(
			fmt.Errorf("外部APIのGETリクエストに対するレスポンス読み込みに失敗しました。url: %s, res.Body: %s", url, res.Body),
		)
		return domain.NewCustomError(
			http.StatusInternalServerError,
			[]string{config.ERROR_MESSAGE_500},
		)
	}

	// ステータスコードチェック
	if res.StatusCode != http.StatusOK {
		infrastructure.NewSentryErrorHandler().SendSentry(
			fmt.Errorf("外部APIのGETリクエストに対するレスポンスのHTTPステータスコードが正常系ではありません。url: %s, httpStatusCode: %d", url, res.StatusCode),
		)
		return domain.NewCustomError(
			http.StatusInternalServerError,
			[]string{config.ERROR_MESSAGE_500},
		)
	}

	// JSONパース
	fmt.Println(string(body))
	if err := json.Unmarshal(body, &response); err != nil {
		infrastructure.NewSentryErrorHandler().SendSentry(
			fmt.Errorf("外部APIのGETリクエストに対するレスポンスのJSONパースに失敗しました。url: %s, body: %s", url, body),
		)
		return domain.NewCustomError(
			http.StatusInternalServerError,
			[]string{config.ERROR_MESSAGE_500},
		)
	}

	return nil
}
