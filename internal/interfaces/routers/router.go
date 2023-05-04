package routers

import (
	"fmt"
	"net/http"
	"news-api/internal/config"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func SetRouters() {
	router := gin.Default()

	// panicをキャッチする
	router.Use(sentryRecovery())

	// Swagger
	setSwaggerRouter(router)

	// NewsAPI
	setNewsAPIRouter(router)

	// RedditAPI (TODO)
	// GoogleNewsAPI (TODO)
	// BingNewsSearchAPI (TODO)

	router.Run(":8080")
}

// panicをキャッチし、Sentryにエラーを送信する
func sentryRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				sentry.CaptureException(fmt.Errorf("panic: %v", r))

				// プログラムが終わる前にSentryがイベントを送信するようにする
				sentry.Flush(time.Second * 2)

				c.JSON(http.StatusInternalServerError, gin.H{
					"error": config.ERROR_MESSAGE_500,
				})
			}
		}()
		// ミドルウェアチェインのために必要
		c.Next()
	}
}
