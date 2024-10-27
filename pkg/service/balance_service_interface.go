package service

type BalanceService interface {
	GetUsableBalance(userID string) (float64, error)
}
