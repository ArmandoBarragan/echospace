package routers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func login(c *gin.Context) {
	var jwtExpirationHours, err = strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOURS"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Configuration error"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "General Kenobi",
		"exp":      time.Now().Add(time.Hour * time.Duration(jwtExpirationHours)).Unix(),
	})
	var secretKey string = os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.AbortWithStatus(500)
	}
	c.IndentedJSON(http.StatusOK, tokenString)
}
