package gateway

import (
	"exinity-task/pkg/model"
	"exinity-task/pkg/types"
	"fmt"

	"github.com/google/uuid"
)

type GatewayB struct{}

func MockSoapApiCall(request types.GatewayBTransactionRequest, c chan types.GatewayBTransactionResponse) {
	response := types.GatewayBTransactionResponse{
		Body: types.GatewayBTransactionResponseBody{
			TransactionResponse: types.GatewayBRawTransactionResponse{
				TransactionID: uuid.New().String(),
				Status:        "SUCCESS",
			},
		},
	}
	c <- response
}

func (g *GatewayB) SendDeposit(transaction model.Transactions) (string, error) {
	requestBody := types.GatewayBTransactionRequest{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Web:   "http://www.example.com/webservices",
		Body: types.SOAPBody{
			Transaction: types.GatewayBRawTransactionRequest{
				Amount: transaction.Amount,
				UserID: transaction.UserID,
				Type:   "CASH_DEPOSIT",
			},
		},
	}

	c := make(chan types.GatewayBTransactionResponse)

	go MockSoapApiCall(requestBody, c)

	response := <-c

	fmt.Println("Transaction ID:", response.Body.TransactionResponse.TransactionID)
	fmt.Println("Status:", response.Body.TransactionResponse.Status)

	return response.Body.TransactionResponse.TransactionID, nil
}

func (g *GatewayB) SendWithdrawal(transaction model.Transactions) (string, error) {
	requestBody := types.GatewayBTransactionRequest{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Web:   "http://www.example.com/webservices",
		Body: types.SOAPBody{
			Transaction: types.GatewayBRawTransactionRequest{
				Amount: transaction.Amount,
				UserID: transaction.UserID,
				Type:   "CASH_WITHDRAW",
			},
		},
	}

	c := make(chan types.GatewayBTransactionResponse)

	go MockSoapApiCall(requestBody, c)

	response := <-c

	fmt.Println("Transaction ID:", response.Body.TransactionResponse.TransactionID)
	fmt.Println("Status:", response.Body.TransactionResponse.Status)

	return response.Body.TransactionResponse.TransactionID, nil
}
