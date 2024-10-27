package service

import "exinity-task/pkg/repository"

type BalanceServiceImpl struct {
	TransactionsRepository repository.TransactionsRepository
}

func NewBalanceService(transactionsRepository repository.TransactionsRepository) BalanceService {
	return &BalanceServiceImpl{TransactionsRepository: transactionsRepository}
}

func (service *BalanceServiceImpl) GetUsableBalance(userID string) (float64, error) {
	// this function is used to calculate the running total of the user balance
	// running total means all completed deposits minus all the withdrawals which are not failed
	// this basically means that the user can only use the funds which are not held or pending

	// fetching all completed deposits
	depositTransactions, err := service.TransactionsRepository.GetAllCompletedDepositsByUserID(userID)

	if err != nil {
		return 0, err
	}

	// fetching all withdrawals which are not failed (meaning they are either pending, held or complete)
	withdrawTransactions, err := service.TransactionsRepository.GetAllNonFailedWithdrawalsByUserID(userID)

	if err != nil {
		return 0, err
	}

	balance := 0.0

	for _, depositTransaction := range depositTransactions {
		balance += depositTransaction.Amount
	}

	for _, withdrawTransaction := range withdrawTransactions {
		balance -= withdrawTransaction.Amount
	}

	return balance, nil

}
