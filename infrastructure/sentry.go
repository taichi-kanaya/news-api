package infrastructure

import (
	"news-api/domain"
	"time"

	"github.com/getsentry/sentry-go"
)

type SentryErrorHandler struct{}

func NewSentryErrorHandler() domain.CustomErrorHandler {
	return &SentryErrorHandler{}
}

// エラーログをSentryに送信する
//
// Parameters:
//   - err: errorインタフェース
//
// Returns:
//   - None
func (seh *SentryErrorHandler) SendSentry(err error) {
	if err != nil {
		sentry.CaptureException(err)

		// プログラムが終わる前にSentryがイベントを送信するようにする
		sentry.Flush(time.Second * 2)
	}
}
