package api

import (
	"stockify/model"

	"github.com/gin-gonic/gin"
)

func SetupAuthenAPI(router *gin.Engine) {
	authenAPIv2 := router.Group("/api/v2")
	{
		authenAPIv2.POST("/login", login)
		authenAPIv2.POST("/register", register)
	}
}

func login(c *gin.Context) {
	c.JSON(201, gin.H{"status": "logged in"})
}

func register(c *gin.Context) {

	var user model.User
	if c.ShouldBind(&user) == nil {
		c.JSON(201, gin.H{"result": "registered", "data": user})
	}

}
