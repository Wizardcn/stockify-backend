package api

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {

	SetupAuthenAPI(router)
	// SetupProductAPI(router)
	// SetupTransactionAPI(router)

}
