package controllers

import (
	"banking_transaction_go/models"
	"time"
)

func NewUserResponse(u *models.User) UserResponse {
	if u == nil {
		return UserResponse{}
	}
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}

func NewAccountResponse(a *models.BankAccount) AccountResponse {
	if a == nil {
		return AccountResponse{}
	}
	return AccountResponse{
		ID:            a.ID,
		AccountNumber: a.AccountNumber,
		Balance:       a.Balance,
		// OwnerID:       a.UserID,
	}
}

func NewTransactionResponse(t *models.Transaction) TransactionResponse {
	if t == nil {
		return TransactionResponse{}
	}
	return TransactionResponse{
		ID:          t.ID,
		ReferenceID: t.ReferenceID,
		Amount:      t.Amount,
		Type:        t.TransactionType,
		CreatedAt:   t.CreatedAt.Format(time.RFC3339),
	}
}

func MapUsers(in []*models.User) []UserResponse {
	out := make([]UserResponse, 0, len(in))
	for _, u := range in {
		out = append(out, NewUserResponse(u))
	}
	return out
}
