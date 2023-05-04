package interceptor

import (
	"fmt"
	"net/http"
	"stockify/model"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey string = "12351935"

func JwtSign(payload model.User) string {
	alClaims := jwt.MapClaims{}
	// Payload begin
	alClaims["id"] = payload.ID
	alClaims["username"] = payload.Username
	alClaims["level"] = payload.Level
	alClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// Payload end
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, alClaims)
	token, _ := at.SignedString(([]byte(secretKey)))
	return token
}

func JwtVertify(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		staffID := fmt.Sprintf("%v", claims["id"])
		username := fmt.Sprintf("%v", claims["username"])
		level := fmt.Sprintf("%v", claims["level"])

		c.Set("jwt_staff_id", staffID)
		c.Set("jwt_username", username)
		c.Set("jwt_level", level)

		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result":  "nok",
			"message": "invalid token",
			"error":   err,
		})
		c.Abort()
	}
}
