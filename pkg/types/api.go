package types

type CreateTransactionRequest struct {
	Amount         float64 `json:"amount" binding:"required,gt=0"`
	UserID         string  `json:"user_id" binding:"required"`
	TypeHandle     string  `json:"type_handle" binding:"required,oneof=DEPOSIT WITHDRAW"`
	ProviderHandle string  `json:"provider_handle" binding:"required"`
}
