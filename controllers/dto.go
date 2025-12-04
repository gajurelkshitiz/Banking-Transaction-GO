package controllers

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

// Response DTOs (optional, can be used as Data payloads)
type AuthPayload struct {
    AccessToken  string      `json:"access_token"`
    RefreshToken string      `json:"refresh_token,omitempty"`
    User         interface{} `json:"user,omitempty"`
}