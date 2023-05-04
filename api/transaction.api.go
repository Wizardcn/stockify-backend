package api

import (
	"net/http"
	"stockify/db"
	"stockify/interceptor"
	"stockify/model"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionResult struct {
	ID            uint
	Total         float64
	Paid          float64
	Change        float64
	PaymentType   string
	PaymentDetail string
	OrderList     string
	Staff         string
	CreatedAt     time.Time
}

func SetupTransactionAPI(router *gin.Engine) {
	transactionAPIv2 := router.Group("api/v2")
	{
		transactionAPIv2.GET("/transaction", getTransactions)
		transactionAPIv2.POST("/transaction", interceptor.JwtVertify, createTransaction)
	}
}

func getTransactions(c *gin.Context) {
	var result []TransactionResult

	db.GetDB().Debug().Raw(`
		SELECT transactions.id, 
			total, 
			paid, 
			change, 
			payment_type, 
			payment_detail,  
			order_list, 
			users.username as staff, 
			transactions.created_at
		FROM transactions JOIN users 
		ON transactions.staff_id = users.id;
	`, nil).Scan(&result)
	c.JSON(http.StatusOK, result)
}

func createTransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBind(&transaction); err == nil {
		transaction.StaffID = c.GetString("jwt_staff_id")
		transaction.CreatedAt = time.Now()
		db.GetDB().Create(&transaction)
		c.JSON(http.StatusCreated, gin.H{"result": "ok", "data": transaction})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"result": "nok"})
	}
}
