package repository

import "exinity-task/pkg/model"

type TransactionsRepository interface {
	Create(transaction *model.Transactions) error
	GetAll() ([]model.Transactions, error)
	Update(transaction *model.Transactions) error
	GetAllCompletedDepositsByUserID(userID string) ([]model.Transactions, error)
	GetAllNonFailedWithdrawalsByUserID(userID string) ([]model.Transactions, error)
	GetTransactionByExternalIDAndGateway(externalID, gateway string) (*model.Transactions, error)
}
