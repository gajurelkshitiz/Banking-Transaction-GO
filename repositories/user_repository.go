package repositories

import (
	"banking_transaction_go/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (r *UserRepository) Create(user models.User) (*models.User, error) {
	err := r.db.Create(&user).Error
	return &user, err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) Exists(email string) bool {
	var count int64
	r.db.Model(&models.User{}).
		Where("email = ?", email).
		Count(&count)
	return count > 0
}
