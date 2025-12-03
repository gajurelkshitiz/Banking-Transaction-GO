package models

import (
	"time"
)

type BankAccount struct {
	ID            uint    `gorm:"primaryKey;autoIncrement"`
	UserID        uint    `gorm:"index"`
	AccountNumber string  `gorm:"type:varchar(20);unique;not null"`
	Balance       float64 `gorm:"type:numeric(15,2);not null;default:0"`
	Status        string  `gorm:"type:varchar(20);not null;default:'active'"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
