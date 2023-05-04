package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"stockify/interceptor"
	"stockify/model"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPIv2 := router.Group("/api/v2")
	{
		productAPIv2.GET("/product", interceptor.JwtVertify, getProduct)
		productAPIv2.POST("/product", createProduct)
	}
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result":   "get product",
		"username": c.GetString("jwt_username"),
		"level":    c.GetString("jwt_level"),
	})
	// c.JSON(http.StatusOK, gin.H{"result": "get product"})
}

func createProduct(c *gin.Context) {
	product := model.Product{}
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	image, _ := c.FormFile("image")
	product.Image = image.Filename

	workingDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/uploaded/images/%s", workingDir, image.Filename)
	c.SaveUploadedFile(image, filePath)

	c.JSON(http.StatusOK, gin.H{"result": product})
}
