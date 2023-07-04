package router

import (
	"saichudin/go-api-gin/product"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.GET("/", product.Home)
	r.GET("/product", product.HandlerList)
	r.GET("/product/:id", product.HandlerDetail)
}