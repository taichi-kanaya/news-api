package domain

type CustomError struct {
	HttpStatusCode int
	Message        string
}

func (e *CustomError) Error() string {
	return e.Message
}

// カスタムエラーを生成する
func NewCustomError(
	httpStatusCode int,
	message string,
) error {
	return &CustomError{
		HttpStatusCode: httpStatusCode,
		Message:        message,
	}
}

// カスタムエラーハンドリング用のインターフェース
type CustomErrorHandler interface {
	SendSentry(err error)
}
