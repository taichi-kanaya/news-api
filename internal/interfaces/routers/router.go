package routers

import "github.com/gin-gonic/gin"

func SetRouters() {
	router := gin.Default()

	// Swagger
	setSwaggerRouter(router)
	// NewsAPI
	setNewsAPIRouter(router)
	// RedditAPI (TODO)
	// GoogleNewsAPI (TODO)
	// BingNewsSearchAPI (TODO)

	router.Run(":8080")
}
