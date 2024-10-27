package repository

import (
	"exinity-task/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.Transactions{})
	return db
}

func TestCreateTransaction(t *testing.T) {
	db := SetupTestDB()
	repo := NewTransactionRepository(db)

	transaction := model.Transactions{
		ExternalID: "txn123",
		UserID:     "user1",
		Amount:     100.0,
		TypeHandle: "DEPOSIT",
	}

	err := repo.Create(&transaction)
	assert.NoError(t, err)
	assert.NotZero(t, transaction.ID)
}

func TestGetTransactionByExternalIDAndGateway(t *testing.T) {
	db := SetupTestDB()
	repo := NewTransactionRepository(db)

	transaction := model.Transactions{
		ExternalID:     "txn123",
		UserID:         "user1",
		Amount:         100.0,
		TypeHandle:     "DEPOSIT",
		ProviderHandle: "gatewayA",
	}

	_ = repo.Create(&transaction)

	foundTransaction, err := repo.GetTransactionByExternalIDAndGateway("txn123", "gatewayA")
	assert.NoError(t, err)
	assert.Equal(t, transaction.ExternalID, foundTransaction.ExternalID)
	assert.Equal(t, transaction.ProviderHandle, foundTransaction.ProviderHandle)
}
