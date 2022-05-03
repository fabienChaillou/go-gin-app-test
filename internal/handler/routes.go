package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	// Handle the index route
	r.GET("/", showIndexPage)

	// Handle GET requests at /article/view/some_article_id
	r.GET("/article/view/:article_id", getArticle)
}
