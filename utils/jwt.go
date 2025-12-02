package utils

import (
	"os"
	"strconv"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func init() {
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID uint) (string, error) {
	minutes, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_MINUTES"))
	expiration := time.Now().Add(time.Duration(minutes) * time.Minute)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// return token.SignedString(jwtSecret)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func GenerateRefreshToken(userID uint) (string, error) {
	days := 7  // default for fallback
	if s := os.Getenv("REFRESH_TOKEN_DAYS"); s != "" {
		if d, err := strconv.Atoi(s); err == nil && d > 0 {
			days = d
		}
	}

	expiration := time.Now().Add(time.Duration(days) * 24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}




func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}