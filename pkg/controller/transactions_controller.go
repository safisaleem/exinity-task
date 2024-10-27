package controller

import (
	constants "exinity-task/pkg/contants"
	"exinity-task/pkg/helper"
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
		helper.RespondWithError(ctx, constants.ErrorInvalidJSON)
		return
	}

	if err := controller.TransactionsService.Create(ctx.Request.Context(), incomingTransaction); err != nil {
		helper.RespondWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, incomingTransaction)
}
