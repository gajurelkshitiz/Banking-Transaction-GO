package controllers

import (
	"banking_transaction_go/services"
	"banking_transaction_go/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type BankAccountController struct {
	Service *services.BankAccountService
}


func getUserIDFromHeader(c echo.Context) (uint, error) {
	auth := c.Request().Header.Get("Authorization")
	if auth == "" {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
	}
	if strings.HasPrefix(strings.ToLower(auth), "bearer ") {
		auth = auth[7:]
	}
	claims, err := utils.ValidateToken(auth)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}

func (bc *BankAccountController) Create(c echo.Context) error {
	var body struct {
		AccountNumber string `json:"account_number"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}

	userID, err := getUserIDFromHeader(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}

	acct, err := bc.Service.Create(userID, body.AccountNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, acct)
}

func (bc *BankAccountController) Get(c echo.Context) error {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	acct, err := bc.Service.GetByID(uint(id64))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, acct)
}

func (bc *BankAccountController) ListByUser(c echo.Context) error {
	userID, err := getUserIDFromHeader(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}
	accts, err := bc.Service.ListByUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, accts)
}

func (bc *BankAccountController) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	if err := bc.Service.Delete(uint(id64)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (bc *BankAccountController) Deposit(c echo.Context) error {
	idStr := c.Param("id")
	id64, _ := strconv.ParseUint(idStr, 10, 64)
	var body struct {
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}
	txn, err := bc.Service.Deposit(uint(id64), body.Amount, body.Description)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, txn)
}

func (bc *BankAccountController) Withdraw(c echo.Context) error {
	idStr := c.Param("id")
	id64, _ := strconv.ParseUint(idStr, 10, 64)
	var body struct {
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}
	txn, err := bc.Service.Withdraw(uint(id64), body.Amount, body.Description)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, txn)
}

func (bc *BankAccountController) Transfer(c echo.Context) error {
	var body struct {
		FromAccountID string  `json:"from_account_no"`
		ToAccountID   string  `json:"to_account_no"`
		Amount        float64 `json:"amount"`
	}

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}

	txn, err := bc.Service.Transfer(body.FromAccountID, body.ToAccountID, body.Amount)
	if err != nil {
		// map service errors to appropriate HTTP status
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"from_account_no": body.FromAccountID,
		"to_account_no":   body.ToAccountID,
		"amount":          body.Amount,
		"transaction_id":  txn.ReferenceID,
	})
}
