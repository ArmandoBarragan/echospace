package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Id      int    `json: "id"`
	Author  string `json: "author"`
	Content string `json: "content"`
}

func homePage(c *gin.Context) {
	var posts = []Post{
		{Id: 1, Author: "Armando", Content: "Hello, there!"},
		{Id: 2, Author: "Anais", Content: "General Kenobi!"},
	}
	c.IndentedJSON(http.StatusOK, posts)
}

func main() {
	var router *gin.Engine = gin.Default()
	router.GET("/home", homePage)
	router.Run("localhost:8000")
}
