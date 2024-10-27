package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitTransactionRoutes(router *gin.Engine) {
	transactionRouter := router.Group("/transactions")

	transactionRouter.POST("", func(context *gin.Context) {
		fmt.Println("here")
		context.JSON(http.StatusOK, "here")
	})
}
