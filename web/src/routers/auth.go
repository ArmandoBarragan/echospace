package routers

import (
	"echospace/src/schemas"
	"echospace/src/utils"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func createUser(c *gin.Context) {
	// Validate user account data and create the account if it passes every validation.
	var newAccountSchema schemas.CreateAccount
	if err := c.ShouldBindJSON(&newAccountSchema); err != nil {
		utils.JSONError(c, 400, err)
		return
	}

	usernameExists, err := newAccountSchema.UsernameExists()

	if err != nil {
		utils.JSONError(c, 400, err)
		return
	}

	if usernameExists {
		utils.JSONError(c, 400, errors.New("that username is already taken"))
		return
	}

	if newAccountSchema.Password != newAccountSchema.PasswordConfirmation {
		utils.JSONError(c, 400, errors.New("passwords don't match"))
		return
	}

	if !newAccountSchema.PasswordIsValid() {
		utils.JSONError(c, 400, errors.New("password is not valid. It requires at least one special character and one digit"))
		return
	}

	newAccount, err := newAccountSchema.Create()

	if err != nil {
		log.Println(err.Error())
		utils.JSONError(c, 500, err)
		return
	}

	utils.JSONSuccess(c, 201, "Account created successfully", gin.H{"id": newAccount.Id})
}

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
