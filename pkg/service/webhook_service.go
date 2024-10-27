package service

import (
	"context"
	"errors"
	constants "exinity-task/pkg/contants"
	"exinity-task/pkg/repository"
	"exinity-task/pkg/types"
)

type WebhookServiceImpl struct {
	TransactionRepository repository.TransactionsRepository
}

func NewWebhookService(transactionRepository repository.TransactionsRepository) WebhookService {
	return &WebhookServiceImpl{
		TransactionRepository: transactionRepository,
	}
}

func (service *WebhookServiceImpl) ProcessWebhook(context context.Context, webhookTransactionEvent types.WebhookTransactionEvent, gatewayType string) error {

	transaction, err := service.TransactionRepository.GetTransactionByExternalIDAndGateway(webhookTransactionEvent.TransactionID, gatewayType)

	if err != nil {
		return err
	}

	if transaction == nil {
		return errors.New("transaction not found")
	}

	if transaction.StatusHandle == constants.TRANSACTION_STATUS_COMPLETE {
		return errors.New("transaction already complete")
	}

	transaction.StatusHandle = webhookTransactionEvent.Status

	err = service.TransactionRepository.Update(transaction)

	if err != nil {
		return err
	}

	return nil
}
