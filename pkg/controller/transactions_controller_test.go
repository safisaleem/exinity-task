package controller

import (
	"bytes"
	"context"
	"exinity-task/pkg/model"
	"exinity-task/pkg/types"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionsService struct {
	mock.Mock
}

func (m *MockTransactionsService) Create(ctx context.Context, transaction types.CreateTransactionRequest) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}

func (m *MockTransactionsService) RetryHeldTransaction(ctx context.Context, transaction *model.Transactions) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}

func (m *MockTransactionsService) GetAllHeldWithdrawals(userID string) ([]model.Transactions, error) {
	args := m.Called(userID)
	return args.Get(0).([]model.Transactions), args.Error(1)
}

func (m *MockTransactionsService) GetAll() ([]model.Transactions, error) {
	args := m.Called()
	return args.Get(0).([]model.Transactions), args.Error(1)
}

func (m *MockTransactionsService) Update(transaction *model.Transactions) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func setupTransactionsRouter() (*gin.Engine, *MockTransactionsService) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := new(MockTransactionsService)
	transactionsController := NewTransactionController(mockService)

	router.POST("/transactions", transactionsController.Create)
	return router, mockService
}

func TestCreateTransaction_Success(t *testing.T) {
	router, mockService := setupTransactionsRouter()

	mockService.On("Create", mock.Anything, mock.Anything).Return(nil)

	payload := `{"amount": 100.0, "user_id": "user1", "type_handle": "DEPOSIT", "provider_handle": "gatewayA"}`
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertCalled(t, "Create", mock.Anything, mock.Anything)
}

func TestCreateTransaction_InvalidJSON(t *testing.T) {
	router, _ := setupTransactionsRouter()

	payload := `{"amount": "invalid_amount", "user_id": "user1", "type_handle": "DEPOSIT", "provider_handle": "gatewayA"}`
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateTransaction_ServiceError(t *testing.T) {
	router, mockService := setupTransactionsRouter()

	mockService.On("Create", mock.Anything, mock.Anything).Return("error insufficient balance")

	payload := `{"amount": 1000.0, "user_id": "user1", "type_handle": "WITHDRAW", "provider_handle": "gatewayA"}`
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
