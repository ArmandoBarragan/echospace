package main

import (
	"fmt"
	"net/http"
	"os"

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

	var host string = os.Getenv("HOST")
	var port string = os.Getenv("PORT")
	router.Run(fmt.Sprintf("%s:%s", host, port))
}
