package main

import (
	"os"
	"log"
	"net/http"

	"banking_transaction_go/database"

	"github.com/labstack/echo/v4"
)

func main() {
	// connect to DB
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}

	// run migrations and handle error
	if err := database.Migrate(db); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	// start server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	e.Logger.Fatal(e.Start(":" + port))
}