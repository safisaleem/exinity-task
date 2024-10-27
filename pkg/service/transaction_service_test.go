package service

import (
	"context"
	"testing"

	constants "exinity-task/pkg/contants"
	"exinity-task/pkg/gateway"
	"exinity-task/pkg/model"
	"exinity-task/pkg/repository"
	"exinity-task/pkg/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MockBalanceService struct {
	mock.Mock
}

func (m *MockBalanceService) GetUsableBalance(userID string) (float64, error) {
	args := m.Called(userID)
	return args.Get(0).(float64), args.Error(1)
}

type MockPaymentGatewayFactory struct {
	mock.Mock
}

func (m *MockPaymentGatewayFactory) GetPaymentGateway(providerHandle string) gateway.PaymentGateway {
	args := m.Called(providerHandle)
	return args.Get(0).(gateway.PaymentGateway)
}

type MockPaymentGateway struct {
	mock.Mock
}

func (m *MockPaymentGateway) SendDeposit(transaction model.Transactions) (string, error) {
	args := m.Called(transaction)
	return args.String(0), args.Error(1)
}

func (m *MockPaymentGateway) SendWithdrawal(transaction model.Transactions) (string, error) {
	args := m.Called(transaction)
	return args.String(0), args.Error(1)
}

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the in-memory database")
	}
	db.AutoMigrate(&model.Transactions{})
	return db
}

func TestCreateTransaction_WithdrawalWithSufficientBalance(t *testing.T) {
	db := SetupTestDB()
	repo := repository.NewTransactionRepository(db)

	mockBalance := new(MockBalanceService)
	mockBalance.On("GetUsableBalance", "user1").Return(200.0, nil)

	mockGateway := new(MockPaymentGateway)
	mockGateway.On("SendWithdrawal", mock.Anything).Return("txn123", nil)

	mockGatewayFactory := new(MockPaymentGatewayFactory)
	mockGatewayFactory.On("GetPaymentGateway", "gatewayA").Return(mockGateway)

	transactionService := NewTransactionService(repo, mockGatewayFactory, mockBalance)

	req := types.CreateTransactionRequest{
		Amount:         100.0,
		UserID:         "user1",
		TypeHandle:     constants.TRANSACTION_TYPE_WITHDRAW,
		ProviderHandle: "gatewayA",
	}

	err := transactionService.Create(context.Background(), req)
	assert.NoError(t, err)
	mockBalance.AssertCalled(t, "GetUsableBalance", "user1")
	mockGateway.AssertCalled(t, "SendWithdrawal", mock.Anything)
}

func TestCreateTransaction_WithdrawalWithInsufficientBalance(t *testing.T) {
	db := SetupTestDB()
	repo := repository.NewTransactionRepository(db)

	mockBalance := new(MockBalanceService)
	mockBalance.On("GetUsableBalance", "user1").Return(50.0, nil)

	mockGatewayFactory := new(MockPaymentGatewayFactory)

	transactionService := NewTransactionService(repo, mockGatewayFactory, mockBalance)

	req := types.CreateTransactionRequest{
		Amount:         100.0,
		UserID:         "user1",
		TypeHandle:     constants.TRANSACTION_TYPE_WITHDRAW,
		ProviderHandle: "gatewayA",
	}

	err := transactionService.Create(context.Background(), req)
	assert.EqualError(t, err, "insufficient balance")
	mockBalance.AssertCalled(t, "GetUsableBalance", "user1")
}
