package config

import "os"

func GetAppEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		panic("Not found APP_ENV")
	}
	return env
}

func GetNewsAPIKey() string {
	newsApiKey := os.Getenv("NEWS_API_KEY")
	if newsApiKey == "" {
		panic("Not found NEWS_API_KEY")
	}
	return newsApiKey
}

func GetSentryDsn() string {
	sentryDsn := os.Getenv("SENTRY_DSN")
	if sentryDsn == "" {
		panic("Not found SENTRY_DSN")
	}
	return sentryDsn
}
