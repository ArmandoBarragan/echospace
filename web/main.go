package main

import (
	"fmt"
	"livaf/conf"
	"livaf/src/routers"

	// "log"

	"github.com/gin-gonic/gin"
)

func main() {
	// neoEngine, err := conf.SetupNeo()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// neoEngine
	var router gin.Engine = *routers.InitRouter()
	router.Run(fmt.Sprintf("%s:%d", conf.Config.Host, conf.Config.Port))
}
