package services

import (
	"errors"
	"banking_transaction_go/models"
	"banking_transaction_go/repositories"
	"banking_transaction_go/utils"
)

type AuthService struct {
	UserRepo repositories.UserRepository
}

func (s *AuthService) Register(name, email, password string) (*models.User, error) {
	if s.UserRepo.Exists(email) {
		return nil, errors.New("email already registered")
	}

	hash, _ := utils.HashPassword(password)

	user := models.User{
		Name: name, 
		Email: email,
		PasswordHash: hash,
	}

	return s.UserRepo.Create(user)
}


func (s *AuthService) Login(email, password string) (string, string, *models.User, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", "", nil, errors.New("invalid email or password")
	}

	if !utils.CheckPassword(password, user.PasswordHash) {
		return "", "", nil, errors.New("invalid email or password")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", nil, err
	}

	return accessToken, refreshToken, user, nil
}