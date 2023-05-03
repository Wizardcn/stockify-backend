package api

import (
	"stockify/db"

	"github.com/gin-gonic/gin"
)

// SetupRouter - invoke this method to setup routes
func SetupRouter(router *gin.Engine) {

	db.SetupDB()

	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupTransactionAPI(router)

}
