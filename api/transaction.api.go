package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupTransactionAPI(router *gin.Engine) {
	transactionAPIv2 := router.Group("api/v2")
	{
		transactionAPIv2.GET("/transaction", getTransaction)
		transactionAPIv2.POST("/transaction", createTransaction)
	}
}

func getTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "get transaction"})
}

func createTransaction(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{"result": "create transaction"})

}
