package service

import (
	"context"
	"exinity-task/pkg/model"
	"exinity-task/pkg/types"
)

type TransactionsService interface {
	Create(context context.Context, transaction types.CreateTransactionRequest) error
	GetAll() ([]model.Transactions, error)
	Update(transaction *model.Transactions) error
}
