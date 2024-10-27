package controller

import (
	"exinity-task/pkg/service"
	"exinity-task/pkg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionsController struct {
	TransactionsService service.TransactionsService
}

func NewTransactionController(transactionService service.TransactionsService) *TransactionsController {
	return &TransactionsController{TransactionsService: transactionService}
}

func (controller *TransactionsController) Create(ctx *gin.Context) {
	var incomingTransaction types.CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&incomingTransaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.TransactionsService.Create(ctx.Request.Context(), incomingTransaction); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, incomingTransaction)
}
