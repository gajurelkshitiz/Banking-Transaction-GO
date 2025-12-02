package database

import (
	"banking_transaction_go/models"

	"gorm.io/gorm"
)

// createEnums ensures Postgres enum types exist.
func createEnums(db *gorm.DB) error {
	if err := db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_type_enum') THEN
				CREATE TYPE transaction_type_enum AS ENUM ('transfer', 'deposit', 'withdraw');
			END IF;
		END$$;
	`).Error; err != nil {
		return err
	}

	if err := db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_status_enum') THEN
				CREATE TYPE transaction_status_enum AS ENUM('pending', 'success', 'failed');
			END IF;
		END$$;
	`).Error; err != nil {
		return err
	}

	return nil
}

// Migrate runs AutoMigrate for your models and returns an error if any.
func Migrate(db *gorm.DB) error {
	if err := createEnums(db); err != nil {
		return err
	}

	return db.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.BankAccount{},
		&models.Transaction{},
	)
}