package routers

import "github.com/gin-gonic/gin"

func jsonSuccess(c *gin.Context, statusCode int, message string, data gin.H) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

func jsonError(c *gin.Context, statusCode int, err error) {
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
			"data":    gin.H{"error": err},
		})
	}
	c.JSON(statusCode, gin.H{
		"message": err.Error(),
	})
}
