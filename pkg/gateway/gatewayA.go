package gateway

import (
	"encoding/json"
	"exinity-task/pkg/model"
	"exinity-task/pkg/types"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type GatewayA struct{}

func MockRestApiCall(url string, jsonData []byte, c chan types.GatewayATransactionResponse) {
	fmt.Println("Making API call to gatewayA")

	var response types.GatewayATransactionResponse

	response.TransactionID = uuid.New().String()
	response.Status = "SUCCESS"

	c <- response

}

func (g *GatewayA) SendDeposit(transaction model.Transactions) (string, error) {
	url := "https://jsonplaceholder.typicode.com/posts"

	jsonData, err := json.Marshal(types.GatewayATransactionRequest{
		Amount: transaction.Amount,
		UserID: transaction.UserID,
		Type:   "CLIENT_DEPOSIT",
	})

	if err != nil {
		log.Fatal(err)
	}

	c := make(chan types.GatewayATransactionResponse)

	go MockRestApiCall(url, jsonData, c)

	response := <-c

	fmt.Println("Transaction ID:", response.TransactionID)
	fmt.Println("Status:", response.Status)

	return response.TransactionID, nil
}

func (g *GatewayA) SendWithdrawal(transaction model.Transactions) (string, error) {
	url := "https://jsonplaceholder.typicode.com/posts"

	jsonData, err := json.Marshal(types.GatewayATransactionRequest{
		Amount: transaction.Amount,
		UserID: transaction.UserID,
		Type:   "CLIENT_WITHDRAW",
	})

	if err != nil {
		log.Fatal(err)
	}

	c := make(chan types.GatewayATransactionResponse)

	go MockRestApiCall(url, jsonData, c)

	response := <-c

	fmt.Println("Transaction ID:", response.TransactionID)
	fmt.Println("Status:", response.Status)

	return response.TransactionID, nil
}
