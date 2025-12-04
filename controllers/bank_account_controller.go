package controllers

import (
	"banking_transaction_go/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BankAccountController struct {
	Service *services.BankAccountService
}

func (bc *BankAccountController) Create(c echo.Context) error {
	// var body struct {
	// 	AccountNumber string `json:"account_number"`
	// }
	var body CreateAccountRequest
	if err := c.Bind(&body); err != nil {
		return JSONError(c, http.StatusBadRequest, "invalid payload", "")
	}

	user_id, err := strconv.Atoi(c.Get("user_id").(string))
	if err != nil {
		return JSONError(c, http.StatusUnauthorized, "userID not found in context", "")
	}
	userID := uint(user_id)

	acct, err := bc.Service.Create(userID, body.AccountNumber)
	if err != nil {
		return JSONError(c, http.StatusBadRequest, err.Error(), "")
	}
	return JSONSuccess(c, http.StatusCreated, acct, "")
}

func (bc *BankAccountController) Get(c echo.Context) error {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return JSONError(c, http.StatusBadRequest, "invalid id", "")
	}
	acct, err := bc.Service.GetByID(uint(id64))
	if err != nil {
		return JSONError(c, http.StatusNotFound, err.Error(), "")
	}
	return JSONSuccess(c, http.StatusOK, acct, "")
}

func (bc *BankAccountController) ListByUser(c echo.Context) error {
	user_id, err := strconv.Atoi(c.Get("user_id").(string))
	if err != nil {
		return JSONError(c, http.StatusUnauthorized, "userID not found in context", "")
	}
	userID := uint(user_id)

	accts, err := bc.Service.ListByUser(userID)
	if err != nil {
		return JSONError(c, http.StatusInternalServerError, err.Error(), "")
	}
	return JSONSuccess(c, http.StatusOK, accts, "")
}

func (bc *BankAccountController) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return JSONError(c, http.StatusBadRequest, "invalid id", "")
	}
	if err := bc.Service.Delete(uint(id64)); err != nil {
		return JSONError(c, http.StatusInternalServerError, err.Error(), "")
	}
	return c.NoContent(http.StatusNoContent)
}

func (bc *BankAccountController) Deposit(c echo.Context) error {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return JSONError(c, http.StatusBadRequest, "invalid id", "")
	}
	var body struct {
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}
	if err := c.Bind(&body); err != nil {
		return JSONError(c, http.StatusBadRequest, "invalid payload", "")
	}
	txn, err := bc.Service.Deposit(uint(id64), body.Amount, body.Description)
	if err != nil {
		return JSONError(c, http.StatusBadRequest, err.Error(), "")
	}
	return JSONSuccess(c, http.StatusOK, txn, "")
}

func (bc *BankAccountController) Withdraw(c echo.Context) error {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return JSONError(c, http.StatusBadRequest, "invalid id", "")
	}
	var body struct {
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}
	if err := c.Bind(&body); err != nil {
		return JSONError(c, http.StatusBadRequest, "invalid payload", "")
	}
	txn, err := bc.Service.Withdraw(uint(id64), body.Amount, body.Description)
	if err != nil {
		return JSONError(c, http.StatusBadRequest, err.Error(), "")
	}
	return JSONSuccess(c, http.StatusOK, txn, "")
}

func (bc *BankAccountController) Transfer(c echo.Context) error {
	var body struct {
		FromAccountID string  `json:"from_account_no"`
		ToAccountID   string  `json:"to_account_no"`
		Amount        float64 `json:"amount"`
	}

	if err := c.Bind(&body); err != nil {
		return JSONError(c, http.StatusBadRequest, "invalid payload", "")
	}

	txn, err := bc.Service.Transfer(body.FromAccountID, body.ToAccountID, body.Amount)
	if err != nil {
		return JSONError(c, http.StatusBadRequest, err.Error(), "")
	}

	return JSONSuccess(c, http.StatusOK, map[string]interface{}{
		"from_account_no": body.FromAccountID,
		"to_account_no":   body.ToAccountID,
		"amount":          body.Amount,
		"transaction_id":  txn.ReferenceID,
	}, "")
}
