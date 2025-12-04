package services

import (
	"banking_transaction_go/models"
	"banking_transaction_go/repositories"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BankAccountService struct {
	BankRepo *repositories.BankAccountRepository
	TxnRepo  *repositories.TransactionRepository
}

func NewBankAccountService(br *repositories.BankAccountRepository, tr *repositories.TransactionRepository) *BankAccountService {
	return &BankAccountService{BankRepo: br, TxnRepo: tr}
}

func (s *BankAccountService) Create(userID uint, accountNumber string) (*models.BankAccount, error) {
	if _, err := s.BankRepo.FindByAccountNumber(accountNumber); err == nil {
		return nil, errors.New("account number already exists")
	}

	acct := models.BankAccount{
		// UserID:        userID,
		AccountNumber: accountNumber,
		Balance:       0,
		Status:        "active",
	}
	return s.BankRepo.Create(acct)
}

func (s *BankAccountService) GetByID(id uint) (*models.BankAccount, error) {
	return s.BankRepo.FindByID(id)
}

func (s *BankAccountService) ListByUser(userID uint) ([]models.BankAccount, error) {
	return s.BankRepo.ListByUser(userID)
}

func (s *BankAccountService) Delete(id uint) error {
	return s.BankRepo.Delete(id)
}

// Deposit increases balance and logs a Transaction atomically.
func (s *BankAccountService) Deposit(accountID uint, amount float64, description string) (*models.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}

	var txn *models.Transaction
	err := s.BankRepo.WithTx(func(tx *gorm.DB) error {
		// change balance inside tx
		if err := s.BankRepo.ChangeBalanceTx(tx, accountID, amount); err != nil {
			return err
		}
		// create txn record using transaction-aware repo
		txn = &models.Transaction{
			ReferenceID:     fmt.Sprintf("txn_%d", time.Now().UnixNano()),
			FromAccountID:   0,
			ToAccountID:     accountID,
			Amount:          amount,
			TransactionType: "deposit",
			Status:          "success",
			Description:     description,
		}
		return s.TxnRepo.CreateWithTx(tx, txn)
	})
	if err != nil {
		return nil, err
	}
	return txn, nil
}

// Decreses balance
func (s *BankAccountService) Withdraw(accountID uint, amount float64, description string) (*models.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}

	var txn *models.Transaction
	err := s.BankRepo.WithTx(func(tx *gorm.DB) error {
		// load account within tx
		acct, err := s.BankRepo.FindByIDTx(tx, accountID)
		if err != nil {
			return err
		}
		if acct.Balance < amount {
			return errors.New("insufficient funds")
		}

		// subtract balance
		if err := s.BankRepo.ChangeBalanceTx(tx, accountID, -amount); err != nil {
			return err
		}

		// create txn record
		txn = &models.Transaction{
			ReferenceID:     fmt.Sprintf("txn_%d", time.Now().UnixNano()),
			FromAccountID:   accountID,
			ToAccountID:     0,
			Amount:          amount,
			TransactionType: "withdraw",
			Status:          "success",
			Description:     description,
		}
		return s.TxnRepo.CreateWithTx(tx, txn)
	})
	if err != nil {
		return nil, err
	}
	return txn, nil
}

func (s *BankAccountService) Transfer(fromAccountNo, toAccountNo string, amount float64) (*models.Transaction, error) {
	if fromAccountNo == toAccountNo {
		return nil, errors.New("cannot transfer to same account")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be > 0")
	}

	var txn *models.Transaction
	err := s.BankRepo.WithTx(func(tx *gorm.DB) error {
		fromAcct, err := s.BankRepo.FindByAccountNumberTx(tx, fromAccountNo)
		if err != nil {
			return err
		}
		toAcct, err := s.BankRepo.FindByAccountNumberTx(tx, toAccountNo)
		if err != nil {
			return err
		}
		if fromAcct.Balance < amount {
			return errors.New("insufficient funds")
		}

		if err := s.BankRepo.ChangeBalanceTx(tx, fromAcct.ID, -amount); err != nil {
			return err
		}
		if err := s.BankRepo.ChangeBalanceTx(tx, toAcct.ID, amount); err != nil {
			return err
		}

		txn = &models.Transaction{
			ReferenceID:     fmt.Sprintf("txn_%d", time.Now().UnixNano()),
			FromAccountID:   fromAcct.ID,
			ToAccountID:     toAcct.ID,
			Amount:          amount,
			TransactionType: "transfer",
			Status:          "success",
		}
		return s.TxnRepo.CreateWithTx(tx, txn)
	})
	if err != nil {
		return nil, err
	}
	return txn, nil
}
