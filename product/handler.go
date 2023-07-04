package product

import (
	"net/http"
	"saichudin/go-api-gin/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id		int64		`json:"id"`
	Name	string		`json:"name"`
	Sku		string		`json:"sku"`
	Price	int64		`json:"price"`
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{"message": "this is home"})
}

func HandlerList(c *gin.Context) {
	products := GetProducts()

	resp := helper.DefaultResponse{
		Message: "success",
		Data: products,
	}

	c.JSON(http.StatusOK, &resp)
}

func HandlerDetail(c *gin.Context) {

	id, _:= strconv.ParseInt(c.Param("id"),0,64)
	product, err:= GetProductDetail(id)

	if err != nil {
		resp := helper.DefaultResponse{
			Message: "success",
			Data: nil,
		}

		c.JSON(http.StatusOK, &resp)
		return
	}

	resp := helper.DefaultResponse{
		Message: "success",
		Data: product,
	}

	c.JSON(http.StatusOK, &resp)
}