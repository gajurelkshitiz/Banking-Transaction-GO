package repositories

import (
	"banking_transaction_go/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) Create(txn *models.Transaction) error {
	return r.db.Create(txn).Error
}
