package main

import (
	"stockify/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/images", "./uploaded/images")

	api.SetupRouter(router)
	router.Run(":8082")
}
