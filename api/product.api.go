package api

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"stockify/db"
	"stockify/model"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPIv2 := router.Group("/api/v2")
	{
		productAPIv2.GET("/product" /*interceptor.JwtVertify,*/, getProduct)
		productAPIv2.POST("/product" /*interceptor.JwtVertify,*/, createProduct)
		productAPIv2.PUT("/product" /*interceptor.JwtVertify,*/, editProduct)
	}
}

func getProduct(c *gin.Context) {

	product := model.Product{}

	id := c.Query("id")
	db.GetDB().First(&product, id)

	c.JSON(http.StatusOK, gin.H{"result": product})
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func saveImage(image *multipart.FileHeader, product *model.Product, c *gin.Context) {
	if image != nil {
		workingDir, _ := os.Getwd()
		product.Image = image.Filename
		extension := filepath.Ext(image.Filename)
		fileName := fmt.Sprintf("%d%s", product.ID, extension)
		filePath := fmt.Sprintf("%s/uploaded/images/%s", workingDir, fileName)

		if fileExists(filePath) {
			os.Remove(filePath)
		}

		c.SaveUploadedFile(image, filePath)

		// Update image path that store on local storage to db
		db.GetDB().Model(&product).Update("image", fileName)
	}
}

func createProduct(c *gin.Context) {
	product := model.Product{}
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

	// Create product on db
	product.CreatedAt = time.Now()
	db.GetDB().Create(&product)

	// Upload product image to local storage
	image, _ := c.FormFile("image")
	saveImage(image, &product, c)

	c.JSON(http.StatusOK, gin.H{"result": product})
}

func editProduct(c *gin.Context) {
	var product model.Product
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 32)
	product.ID = uint(id)
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

	db.GetDB().Save(&product) // pass by ref
	image, _ := c.FormFile("image")
	saveImage(image, &product, c)
	c.JSON(http.StatusOK, gin.H{"result": product})
}
