package interceptor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GeneralInterceptor - call this method to use as general interceptor
func GeneralInterceptor(c *gin.Context) {
	token := c.Query("token")
	if token == "1234" {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		c.Abort()
	}
}
