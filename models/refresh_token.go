package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	ID uint `gorm:"primaryKey"`
	UserID uint	`gorm:"index;not null"`
	Token string `gorm:"not null;uniqueIndex"`
	ExpiresAt time.Time `gorm:"index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// set default expiry from env if not provided
func (rt *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	if rt.ExpiresAt.IsZero() {
		rt.ExpiresAt = time.Now().Add(7 * 24 * time.Hour) // default 7 days
	}
	return nil
}