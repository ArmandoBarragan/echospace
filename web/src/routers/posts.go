package routers

import (
	"echospace/src/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	var posts = []schemas.Post{
		{Id: 1, Author: "Obi Wan Kenobi", Content: "Hello, there!"},
		{Id: 2, Author: "General Grievous", Content: "General Kenobi!"},
	}
	c.IndentedJSON(http.StatusOK, posts)
}
