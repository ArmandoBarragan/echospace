package test

import (
	"echospace/conf"
	"echospace/src/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func testRequest(w http.ResponseWriter, req http.Request) {
	var router gin.Engine = *routers.InitRouter()
	conf.SetupNeo()
	router.ServeHTTP(w, &req)
}
