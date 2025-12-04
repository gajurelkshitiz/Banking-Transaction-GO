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

func createForeignKeys(db *gorm.DB) error {
	// bank_accounts.user_id -> users.id
	if err := db.Exec(`
		DO $$
		BEGIN 
			IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_bank_accounts_user_id') THEN
				ALTER TABLE bank_accounts
				ADD CONSTRAINT fk_bank_accounts_user_id
				FOREIGN KEY (user_id) REFERENCES users(id)
				ON UPDATE CASCADE ON DELETE CASCADE;
			END IF;
		END$$;
	`).Error; err != nil {
		return err
	}

	// transactions.from_account_id -> bank_accounts.id
    if err := db.Exec(`
        DO $$
        BEGIN
          IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_transactions_from_account_id') THEN
            ALTER TABLE transactions
            ADD CONSTRAINT fk_transactions_from_account_id
            FOREIGN KEY (from_account_id) REFERENCES bank_accounts(id)
            ON UPDATE CASCADE ON DELETE RESTRICT;
          END IF;
        END$$;
    `).Error; err != nil {
        return err
    }

    // transactions.to_account_id -> bank_accounts.id
    if err := db.Exec(`
        DO $$
        BEGIN
          IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_transactions_to_account_id') THEN
            ALTER TABLE transactions
            ADD CONSTRAINT fk_transactions_to_account_id
            FOREIGN KEY (to_account_id) REFERENCES bank_accounts(id)
            ON UPDATE CASCADE ON DELETE RESTRICT;
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

	if err := db.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.BankAccount{},
		&models.Transaction{},
	); err != nil {
		return err
	}

	// ensure FK constraints exist on existing DBs
	return createForeignKeys(db)
}
