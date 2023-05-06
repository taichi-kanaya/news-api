package domain

import (
	"fmt"
	"strings"
)

type CustomError struct {
	HttpStatusCode int
	Messages       []string
}

// エラー情報を文字列で返す
//
// Returns:
//   - string: エラー情報
func (e *CustomError) Error() string {
	return fmt.Sprintf("%d:%s", e.HttpStatusCode, strings.Join(e.Messages, ","))
}

// カスタムエラーを生成する
//
// Parameters:
//   - httpStatusCode: HTTPステータスコード
//   - messages: エラーメッセージ
//
// Returns:
//   - error: エラー情報
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
