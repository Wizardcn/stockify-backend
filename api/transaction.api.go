package api

import (
	"net/http"
	"stockify/db"
	"stockify/model"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupTransactionAPI(router *gin.Engine) {
	transactionAPIv2 := router.Group("api/v2")
	{
		transactionAPIv2.GET("/transaction", getTransactions)
		transactionAPIv2.POST("/transaction", createTransaction)
	}
}

func getTransactions(c *gin.Context) {
	var transactions []model.Transaction
	db.GetDB().Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

func createTransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBind(&transaction); err == nil {
		transaction.CreatedAt = time.Now()
		db.GetDB().Create(&transaction)
		c.JSON(http.StatusCreated, gin.H{"result": "ok", "data": transaction})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"result": "nok"})
	}

}
