package models

import "time"

type User struct {
	ID 				uint `gorm:"primaryKey"`
	Name 			string
	Email 			string `gorm:"unique; not null"`
	PasswordHash 	string
	CreatedAt 		time.Time `gorm:"autoCreateTime"`
	UpdatedAt 		time.Time `gorm:"autoUpdateTime"`
}