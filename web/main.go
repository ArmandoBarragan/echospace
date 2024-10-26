package main

import (
	"echospace/conf"
	"echospace/src/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	var router gin.Engine = *routers.InitRouter()
	conf.SetupNeo()
	router.Run(fmt.Sprintf("%s:%d", conf.Config.Host, conf.Config.Port))
	defer conf.ClosePools()
}
