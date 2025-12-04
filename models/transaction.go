package models

import "time"

type Transaction struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	ReferenceID string `gorm:"type:varchar(50);unique;not null"`

	FromAccountID uint `gorm:"not null"`
	ToAccountID uint `gorm:"not null"`

	// Associations
	FromAccount BankAccount `gorm:"foreignKey:FromAccountID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ToAccount BankAccount `gorm:"foreignKey:ToAccountID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	Amount float64 `gorm:"type:numeric(15,2);not null"`

	TransactionType string `gorm:"type:transaction_type_enum;not null;default:'transfer'"` 
	Status string `gorm:"type:transaction_status_enum;not null;default:'pending'"` 

	Description string `gorm:"type:text"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}