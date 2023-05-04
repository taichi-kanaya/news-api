package domain

import "strings"

type CustomError struct {
	HttpStatusCode int
	Messages       []string
}

func (e *CustomError) Error() string {
	return strings.Join(e.Messages, ",")
}

// カスタムエラーを生成する
func NewCustomError(
	httpStatusCode int,
	messages []string,
) error {
	return &CustomError{
		HttpStatusCode: httpStatusCode,
		Messages:       messages,
	}
}

// カスタムエラーハンドリング用のインターフェース
type CustomErrorHandler interface {
	SendSentry(err error)
}
