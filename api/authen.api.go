package api

import (
	"net/http"
	"stockify/db"
	"stockify/model"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	var user model.User

	if c.ShouldBind(&user) == nil {
		var queryUser model.User
		if err := db.GetDB().First(&queryUser, "username = ?", user.Username).Error; err != nil {
			c.JSON(200, gin.H{"result": "nok", "error": err})
		} else if !checkPasswordHash(user.Password, queryUser.Password) {
			c.JSON(200, gin.H{"result": "nok", "error": "invalid password"})
		} else {
			// token := interceptor.JwtSign(queryUser)
			alClaims := jwt.MapClaims{}
			alClaims["id"] = queryUser.ID
			alClaims["username"] = queryUser.Username
			alClaims["level"] = queryUser.Level
			alClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
			at := jwt.NewWithClaims(jwt.SigningMethodHS256, alClaims)
			token, _ := at.SignedString(([]byte("1234")))
			c.JSON(200, gin.H{"result": "ok", "data": token})
		}

	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}
}

func register(c *gin.Context) {

	var user model.User

	if c.ShouldBind(&user) == nil {
		user.Password, _ = hashPassword(user.Password)
		user.CreatedAt = time.Now()
		if err := db.GetDB().Create(&user).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "nok", "error": err})
		} else {
			var queryUser model.User
			db.GetDB().First(&queryUser, "username = ?", user.Username)
			c.JSON(http.StatusOK, gin.H{"result": "ok", "data": queryUser})
		}

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "unable to bind data"})
	}

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
