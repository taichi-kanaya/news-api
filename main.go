package main

import (
	"log"
	"news-api/internal/config"
	"news-api/internal/interfaces/routers"
	"time"

	"github.com/getsentry/sentry-go"
)

// @title			NewsAPI
// @version			1.0
// @license.name	taichi kanaya
// @description		最新のニュースを取得するAPIです
func main() {
	sentryInit()
	routers.SetRouters()
}

// Sentryクライアントの初期設定
func sentryInit() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.GetSentryDsn(),
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)
}
