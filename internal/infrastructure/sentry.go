package infrastructure

import (
	"news-api/internal/domain"

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
	}
}
