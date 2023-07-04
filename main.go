package main

import (
	"saichudin/go-api-gin/config"
	"saichudin/go-api-gin/router"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDb()

	r := gin.Default()
	router.Route(r)
	r.Run()
}
