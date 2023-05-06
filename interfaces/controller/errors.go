package controller

// クライアントに返すエラーレスポンス
type ErrorResponse struct {
	Errors []string `json:"errors"`
}
