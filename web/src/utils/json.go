package utils

import "github.com/gin-gonic/gin"

func JSONSuccess(c *gin.Context, statusCode int, message string, data gin.H) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

func JSONError(c *gin.Context, statusCode int, err error) {
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(statusCode, gin.H{})
	}
}
