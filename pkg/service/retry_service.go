package service

import (
	"context"
	"fmt"
	"time"
)

type RetryServiceImpl struct {
	transactionsService TransactionsService
}

func NewRetryService(transactionsService TransactionsService) RetryService {
	return &RetryServiceImpl{transactionsService: transactionsService}
}

func (service *RetryServiceImpl) RetryHeldTransactions(context context.Context) {

	heldTransactions, err := service.transactionsService.GetAllHeldWithdrawals("")

	if err != nil {
		fmt.Println("Error getting held transactions", err)
		return
	}

	fmt.Println("Retrying", len(heldTransactions), "held transactions")

	for _, heldTransaction := range heldTransactions {

		err := service.transactionsService.RetryHeldTransaction(context, &heldTransaction)

		if err != nil {
			fmt.Println("Error retrying transaction", err)
		}
	}

}

func (service *RetryServiceImpl) Start(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				service.RetryHeldTransactions(ctx)
			case <-ctx.Done():
				fmt.Println("Stopping RetryService")
				return
			}
		}
	}()
}
