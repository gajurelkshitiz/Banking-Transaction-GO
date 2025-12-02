package main

import (
	"os"
	"log"
	"net/http"

	"banking_transaction_go/database"
	"banking_transaction_go/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// connect to DB
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	// repo:=repositories.NewUserRepo(db)
	// service:=service.NewServiceRepo(repo)

	// run migrations and handle error
	if err := database.Migrate(db); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	// start server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routes.AuthRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	e.Logger.Fatal(e.Start(":" + port))
}