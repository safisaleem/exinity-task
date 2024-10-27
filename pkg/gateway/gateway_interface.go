package gateway

import "exinity-task/pkg/model"

type PaymentGateway interface {
	SendDeposit(transaction model.Transactions) (string, error)
	SendWithdrawal(transaction model.Transactions) (string, error)
}
