package repositories

import (
	"banking_transaction_go/models"

	"gorm.io/gorm"
)

type BankAccountRepository struct {
	db *gorm.DB
}

func NewBankAccountRepository(db *gorm.DB) *BankAccountRepository {
	return &BankAccountRepository{db}
}

func (r *BankAccountRepository) Begin() *gorm.DB {
	return r.db.Begin()
}

func (r *BankAccountRepository) Create(account models.BankAccount) (*models.BankAccount, error) {
	err := r.db.Create(&account).Error
	return &account, err
}

func (r *BankAccountRepository) FindByID(id uint) (*models.BankAccount, error) {
	var acct models.BankAccount
	err := r.db.First(&acct, id).Error
	return &acct, err
}

func (r *BankAccountRepository) FindByAccountNumber(number string) (*models.BankAccount, error) {
	var acct models.BankAccount
	err := r.db.Where("account_number = ?", number).First(&acct).Error
	return &acct, err
}

func (r *BankAccountRepository) FindByIDTx(tx *gorm.DB, id uint) (*models.BankAccount, error) {
	var acct models.BankAccount
	if err := tx.First(&acct, id).Error; err != nil {
		return nil, err
	}
	return &acct, nil
}

func (r *BankAccountRepository) FindByAccountNumberTx(tx *gorm.DB, number string) (*models.BankAccount, error) {
	var acct models.BankAccount
	if err := tx.Where("account_number = ?", number).First(&acct).Error; err != nil {
		return nil, err
	}
	return &acct, nil
}

func (r *BankAccountRepository) Delete(id uint) error {
	return r.db.Delete(&models.BankAccount{}, id).Error
}

func (r *BankAccountRepository) ListByUser(userID uint) ([]models.BankAccount, error) {
	var out []models.BankAccount
	err := r.db.Where("user_id = ?", userID).Find(&out).Error
	return out, err
}

func (r *BankAccountRepository) Update(acct *models.BankAccount) error {
	return r.db.Save(acct).Error
}

func (r *BankAccountRepository) ChangeBalance(id uint, delta float64) error {
	return r.db.Model(&models.BankAccount{}).
		Where("id = ?", id).
		UpdateColumn("balance", gorm.Expr("balance + ?", delta)).Error
}

func (r *BankAccountRepository) ChangeBalanceTx(tx *gorm.DB, id uint, delta float64) error {
	return tx.Model(&models.BankAccount{}).
		Where("id = ?", id).
		UpdateColumn("balance", gorm.Expr("balance + ?", delta)).Error
}

// WithTx executes the provided function inside a DB transaction.
// If fn returns an error the transaction is rolled back; otherwise committed.
// It also recovers from panics to rollback safely.
func (r *BankAccountRepository) WithTx(fn func(tx *gorm.DB) error) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// recover from panic and rollback
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
