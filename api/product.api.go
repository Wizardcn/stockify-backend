package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func myInterceptor(c *gin.Context) {
	token := c.Query("token")
	if token == "1234" {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		c.Abort()
	}

}

func SetupProductAPI(router *gin.Engine) {
	productAPIv2 := router.Group("/api/v2")
	{
		productAPIv2.GET("/product", myInterceptor, getProduct)
		productAPIv2.POST("/product", myInterceptor, createProduct)
	}
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "get product"})
}

func createProduct(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"result": "create product"})
}
