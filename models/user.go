package models

import "time"

type User struct {
	ID           uint `gorm:"primaryKey; autoIncrement"`
	Name         string
	Email        string `gorm:"unique; not null"`
	PasswordHash string

	// Relationship : one user -> many bank accounts
	BankAccounts []BankAccount `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
