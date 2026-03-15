package user

import (
	"errors"
	"user-flow/internal/app/models"

	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("user not found")
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Show(id int64) (models.User, error) {
	var user models.User

	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, ErrNotFound
		}

		return models.User{}, err
	}

	return user, nil
}
