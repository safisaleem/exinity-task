package repository

import (
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
