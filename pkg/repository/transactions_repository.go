package repository

import (
	constants "exinity-task/pkg/contants"
	"exinity-task/pkg/model"

	"gorm.io/gorm"
)

type TransactionsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionsRepository {
	return &TransactionsRepositoryImpl{Db: db}
}

func (t *TransactionsRepositoryImpl) Create(transaction *model.Transactions) error {
	return t.Db.Create(&transaction).Error
}

func (t *TransactionsRepositoryImpl) GetAll() ([]model.Transactions, error) {
	var transactions []model.Transactions
	err := t.Db.Find(&transactions).Error
	return transactions, err
}

func (t *TransactionsRepositoryImpl) Update(transaction *model.Transactions) error {
	return t.Db.Save(&transaction).Error
}

func (t *TransactionsRepositoryImpl) GetAllCompletedDepositsByUserID(userID string) ([]model.Transactions, error) {
	var transactions []model.Transactions

	err := t.Db.Where("type_handle = ? AND status_handle = ? AND user_id = ?",
		constants.TRANSACTION_TYPE_DEPOSIT,
		constants.TRANSACTION_STATUS_COMPLETE,
		userID,
	).Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *TransactionsRepositoryImpl) GetAllNonFailedWithdrawalsByUserID(userID string) ([]model.Transactions, error) {
	var transactions []model.Transactions

	err := t.Db.Where("type_handle = ? AND status_handle IN (?, ?, ?) AND user_id = ?",
		constants.TRANSACTION_TYPE_WITHDRAW,
		constants.TRANSACTION_STATUS_PENDING,
		constants.TRANSACTION_STATUS_HELD,
		constants.TRANSACTION_STATUS_COMPLETE,
		userID,
	).Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *TransactionsRepositoryImpl) GetTransactionByExternalIDAndGateway(externalID, gateway string) (*model.Transactions, error) {
	var transaction model.Transactions

	err := t.Db.Where("external_id = ? AND provider_handle = ?", externalID, gateway).First(&transaction).Error

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
