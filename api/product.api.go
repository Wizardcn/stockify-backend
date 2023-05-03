package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPIv2 := router.Group("/api/v2")
	{
		productAPIv2.GET("/product", getProduct)
		productAPIv2.POST("/product", createProduct)
	}
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "get product"})
}

func createProduct(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"result": "create product"})
}
