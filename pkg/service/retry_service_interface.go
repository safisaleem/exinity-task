package service

import "context"

type RetryService interface {
	RetryHeldTransactions(context context.Context)
	Start(context context.Context)
}
