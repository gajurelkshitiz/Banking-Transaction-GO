package controllers

import "time"

// Request DTOs
type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type CreateAccountRequest struct {
	AccountNumber string `json:"account_number" validate:"required,min=6,max=20"`
}

type AmountRequest struct {
	Amount      float64 `json:"amount" validate:"required,gt=0"`
	Description string  `json:"description,omitempty" validate:"omitempty,max=255"`
}

type TransferRequest struct {
	FromAccountID string  `json:"from_account_no" validate:"required"`
	ToAccountID   string  `json:"to_account_no" validate:"required"`
	Amount        float64 `json:"amount" validate:"required,gt=0"`
}

// Response / view DTOs - shape exposed to clients
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	// add only fields you want to expose
}

type AccountResponse struct {
	ID            uint    `json:"id"`
	AccountNumber string  `json:"account_number"`
	Balance       float64 `json:"balance"` // example type
	OwnerID       uint    `json:"owner_id"`
}

type TransactionResponse struct {
	ID          uint    `json:"id"`
	ReferenceID string  `json:"reference_id"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	CreatedAt   string  `json:"created_at"`
}

type AuthPayload struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token,omitempty"`
	User         UserResponse `json:"user"`
}
