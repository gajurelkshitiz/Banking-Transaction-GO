package controllers

import "time"

// Request DTOs
type RegisterRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type RefreshRequest struct {
    RefreshToken string `json:"refresh_token"`
}

type CreateAccountRequest struct {
    AccountNumber string `json:"account_number"`
}

type AmountRequest struct {
    Amount      float64 `json:"amount"`
    Description string  `json:"description,omitempty"`
}

type TransferRequest struct {
    FromAccountID string  `json:"from_account_no"`
    ToAccountID   string  `json:"to_account_no"`
    Amount        float64 `json:"amount"`
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
    ID            uint   `json:"id"`
    AccountNumber string `json:"account_number"`
    Balance       float64  `json:"balance"` // example type
    OwnerID       uint   `json:"owner_id"`
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