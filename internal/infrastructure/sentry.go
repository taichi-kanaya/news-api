package infrastructure

import (
	"news-api/internal/domain"
	"time"

	"github.com/getsentry/sentry-go"
)

type SentryErrorHandler struct{}

func NewSentryErrorHandler() domain.CustomErrorHandler {
	return &SentryErrorHandler{}
}

// エラーログをSentryに送信する
func (seh *SentryErrorHandler) SendSentry(err error) {
	if err != nil {
		sentry.CaptureException(err)

		// プログラムが終わる前にSentryがイベントを送信するようにする
		sentry.Flush(time.Second * 2)
	}
}
