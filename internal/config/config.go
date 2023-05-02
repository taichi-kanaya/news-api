package config

import "os"

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
