package constants

import "errors"

const (
	TRANSACTION_STATUS_PENDING  = "PENDING"
	TRANSACTION_STATUS_HELD     = "HELD"
	TRANSACTION_STATUS_COMPLETE = "COMPLETE"
	TRANSACTION_STATUS_FAILED   = "FAILED"
)

const (
	TRANSACTION_TYPE_DEPOSIT  = "DEPOSIT"
	TRANSACTION_TYPE_WITHDRAW = "WITHDRAW"
)

var ErrorInsufficientBalance = errors.New("insufficient balance")
var ErrorInvalidProvider = errors.New("invalid provider")
var ErrorInvalidTransactionType = errors.New("invalid transaction type")
var ErrorTransactionHeld = errors.New("transaction held")
var ErrorInternal = errors.New("internal error")
var ErrorInvalidJSON = errors.New("invalid json")
