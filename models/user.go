package models

import "time"

type User struct {
	ID             uint `gorm:"primaryKey; autoIncrement"`
	Name           string
	Email          string `gorm:"unique; not null"`
	PasswordHash   string
	BankAccountsID *uint	`gorm:"column:bank_accounts_id;default:NULL"`

	// Relationship : one user ->  bank accounts
	BankAccounts *BankAccount `gorm:"foreignKey:BankAccountsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
