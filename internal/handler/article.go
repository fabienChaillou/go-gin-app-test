package handler

import (
	"net/http"
	"strconv"

	httpinternal "github.com/fabienChaillou/go-gin-app-test/internal/http"
	"github.com/fabienChaillou/go-gin-app-test/internal/model"
	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := model.GetAllArticles()

	// Call the HTML method of the Context to render a template
	httpinternal.Render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := model.GetArticleByID(articleID); err == nil {
			// Call the HTML method of the Context to render a template
			httpinternal.Render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
