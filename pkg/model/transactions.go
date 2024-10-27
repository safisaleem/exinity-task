package model

import "time"

type Transactions struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	ExternalID     string    `gorm:"size:100" json:"external_id"`
	UserID         string    `gorm:"size:50;not null" json:"user_id"`
	Amount         float64   `gorm:"not null" json:"amount"`
	TypeHandle     string    `gorm:"size:50" json:"type_handle"`
	StatusHandle   string    `gorm:"size:50" json:"status_handle"`
	ProviderHandle string    `gorm:"size:50" json:"provider_handle"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
