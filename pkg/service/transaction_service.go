package service

import (
	"context"
	"exinity-task/pkg/model"
	"exinity-task/pkg/repository"
	"exinity-task/pkg/types"
)

type TransactionsServiceImpl struct {
	TransactionsRepository repository.TransactionsRepository
}

func NewTransactionService(transactionRepository repository.TransactionsRepository) TransactionsService {
	return &TransactionsServiceImpl{
		TransactionsRepository: transactionRepository,
	}
}

func (t *TransactionsServiceImpl) Create(context context.Context, incomingTransaction types.CreateTransactionRequest) error {
	transaction := model.Transactions{
		ExternalID:     "",
		Amount:         incomingTransaction.Amount,
		UserID:         incomingTransaction.UserID,
		TypeHandle:     incomingTransaction.TypeHandle,
		ProviderHandle: incomingTransaction.ProviderHandle,
	}

	t.TransactionsRepository.Create(&transaction)

	return nil
}

func (t *TransactionsServiceImpl) GetAll() ([]model.Transactions, error) {
	return t.TransactionsRepository.GetAll()
}

func (t *TransactionsServiceImpl) Update(transaction *model.Transactions) error {
	t.TransactionsRepository.Update(transaction)
	return nil
}
