package repositories

import (
	"banking_transaction_go/database"
	"banking_transaction_go/models"
)

type UserRepository struct{
	db *gorm.DB
}

func NewUserRepo (db *gorm.DB){
	return UserRepository{
		db:db
	}
}
func (UserRepository) Create(user models.User) (*models.User, error) {
	err :=db.Create(&user).Error
	return &user, err
}

func (UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (UserRepository) Exists(email string) bool {
	var count int64
	database.DB.Model(&models.User{}).
		Where("email = ?", email).
		Count(&count)
	return count > 0
}
