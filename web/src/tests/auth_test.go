package test

import (
	"livaf/src/routers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	var router gin.Engine = *routers.InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/create_user", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}
