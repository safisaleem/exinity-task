package helper

import (
	constants "exinity-task/pkg/contants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondWithError(ctx *gin.Context, err error) {
	var statusCode int

	switch err {
	case constants.ErrorInsufficientBalance, constants.ErrorInvalidProvider, constants.ErrorInvalidTransactionType, constants.ErrorInvalidJSON:
		statusCode = http.StatusBadRequest // 400
	case constants.ErrorTransactionHeld:
		statusCode = http.StatusAccepted // 202
	case constants.ErrorInternal:
		statusCode = http.StatusInternalServerError // 500
	default:
		statusCode = http.StatusInternalServerError // 500 for unexpected errors
	}

	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}
