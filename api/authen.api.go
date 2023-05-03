package api

import (
	"net/http"
	"stockify/db"
	"stockify/model"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		user.Password, _ = hashPassword(user.Password)
		user.CreatedAt = time.Now()
		if err := db.GetDB().Create(&user).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "nok", "error": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "ok", "data": user})
		}

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "unable to bind data"})
	}

}

// func checkPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
