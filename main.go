package main

import (
	"log"
	"news-api/config"
	"news-api/infrastructure/routers"
	"strconv"
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
	const appVersion = 1.0
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.GetSentryDsn(),
		TracesSampleRate: 1.0,
		SendDefaultPII:   true,
		Release:          strconv.Itoa(appVersion),
		Environment:      config.GetAppEnv(),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)
}
