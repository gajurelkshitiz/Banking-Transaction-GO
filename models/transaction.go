package models

import "time"

type Transaction struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	ReferenceID string `gorm:"type:varchar(50);unique;not null"`

	FromAccountID uint `gorm:"not null"`
	ToAccountID uint `gorm:"not null"`

	Amount float64 `gorm:"type:numeric(15,2);not null"`

	TransactionType string `gorm:"type:transaction_type_enum;not null;default:'transfer'"` // transfer, deposit, withdraw
	Status string `gorm:"type:transaction_status_enum;not null;default:'pending'"` // pending, success, failed

	Description string `gorm:"type:text"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}