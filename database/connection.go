package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	_ = godotenv.Load()

	// allow single DSN override
	if dsn := strings.TrimSpace(os.Getenv("DATABASE_DSN")); dsn != "" {
		return openDB(dsn)
	}

	required := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT"}
	var missing []string
	for _, k := range required {
		if strings.TrimSpace(os.Getenv(k)) == "" {
			missing = append(missing, k)
		}
	}
	if len(missing) > 0 {
		return nil, fmt.Errorf("missing env vars: %s", strings.Join(missing, ","))
	}

	// Build a dsn from env:
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	return openDB(dsn)
}

func openDB(dsn string) (*gorm.DB, error) {
	// Connect to a postgresql database:
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm DB: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}
	
	DB = db 
	fmt.Println("Database connected successfully!")
	
	return db, nil

}

