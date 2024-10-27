package service

import (
	"context"
	"errors"
	constants "exinity-task/pkg/contants"
	"exinity-task/pkg/gateway"
	"exinity-task/pkg/model"
	"exinity-task/pkg/repository"
	"exinity-task/pkg/types"
	"fmt"
)

type TransactionsServiceImpl struct {
	TransactionsRepository repository.TransactionsRepository
	PaymentGatewayFactory  gateway.PaymentGatewayFactory
	BalanceService         BalanceService
}

func NewTransactionService(transactionRepository repository.TransactionsRepository, gatewayFactory gateway.PaymentGatewayFactory, balanceService BalanceService) TransactionsService {
	return &TransactionsServiceImpl{
		TransactionsRepository: transactionRepository,
		PaymentGatewayFactory:  gatewayFactory,
		BalanceService:         balanceService,
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

	// if transaction type is withdraw
	if incomingTransaction.TypeHandle == constants.TRANSACTION_TYPE_WITHDRAW {
		// check if user has enough usable balance. this ensures that the user cant double spend
		usableBalance, err := t.BalanceService.GetUsableBalance(incomingTransaction.UserID)
		if err != nil {
			return err
		}

		if usableBalance < incomingTransaction.Amount {
			return errors.New("insufficient balance")
		}
	}

	paymentGateway := t.PaymentGatewayFactory.GetPaymentGateway(incomingTransaction.ProviderHandle)

	if paymentGateway == nil {
		transaction.StatusHandle = constants.TRANSACTION_STATUS_FAILED
		t.TransactionsRepository.Create(&transaction)
		return errors.New("invalid provider")
	}

	var externalID string
	var err error

	switch transaction.TypeHandle {
	case constants.TRANSACTION_TYPE_DEPOSIT:
		fmt.Println("sending deposit")
		externalID, err = paymentGateway.SendDeposit(transaction)
	case constants.TRANSACTION_TYPE_WITHDRAW:
		fmt.Println("sending withdraw")
		externalID, err = paymentGateway.SendWithdrawal(transaction)
	default:
		return errors.New("invalid transaction type")
	}

	if err != nil {
		// if the api call to the gateway fails, we hold the transaction
		transaction.StatusHandle = constants.TRANSACTION_STATUS_HELD
		t.TransactionsRepository.Create(&transaction)
		return err
	}

	transaction.ExternalID = externalID
	transaction.StatusHandle = constants.TRANSACTION_STATUS_PENDING

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

func (t *TransactionsServiceImpl) GetAllHeldWithdrawals(userID string) ([]model.Transactions, error) {
	return t.TransactionsRepository.GetAllHeldWithdrawals(userID)
}

func (t *TransactionsServiceImpl) RetryHeldTransaction(context context.Context, incomingTransaction *model.Transactions) error {
	// check if user has enough balance
	usableBalance, err := t.BalanceService.GetUsableBalance(incomingTransaction.UserID)

	fmt.Println("Usable balance", usableBalance)

	if err != nil {
		return err
	}

	// if type is withdrawal
	if incomingTransaction.TypeHandle == constants.TRANSACTION_TYPE_WITHDRAW {
		heldTransactions, err := t.TransactionsRepository.GetAllHeldWithdrawals(incomingTransaction.UserID)

		if err != nil {
			return err
		}

		// we can add the held balance to the usable balance because the held balance is not yet deducted from the usable balance

		heldBalance := 0.0

		for _, heldTransaction := range heldTransactions {
			heldBalance += heldTransaction.Amount
		}

		usableBalance += heldBalance

		// if the user has enough balance, we update the transaction and send it to the gateway

		if usableBalance < incomingTransaction.Amount {
			return errors.New("insufficient balance")
		}
	}

	// now we call the gateway
	paymentGateway := t.PaymentGatewayFactory.GetPaymentGateway(incomingTransaction.ProviderHandle)

	if paymentGateway == nil {
		incomingTransaction.StatusHandle = constants.TRANSACTION_STATUS_FAILED
		t.TransactionsRepository.Update(incomingTransaction)
		return errors.New("invalid provider")
	}

	var externalID string

	switch incomingTransaction.TypeHandle {
	case constants.TRANSACTION_TYPE_DEPOSIT:
		fmt.Println("sending deposit")
		externalID, err = paymentGateway.SendDeposit(*incomingTransaction)
	case constants.TRANSACTION_TYPE_WITHDRAW:
		fmt.Println("sending withdraw")
		externalID, err = paymentGateway.SendWithdrawal(*incomingTransaction)
	default:
		return errors.New("invalid transaction type")
	}

	if err != nil {
		// if the api call to the gateway fails, we keep the transaction as held
		return err
	}

	incomingTransaction.ExternalID = externalID
	incomingTransaction.StatusHandle = constants.TRANSACTION_STATUS_PENDING

	t.TransactionsRepository.Update(incomingTransaction)

	return nil
}
