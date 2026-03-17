package user

import (
	"errors"
	"fmt"
	filter "user-flow/internal/app/domain/core/filter/user"
	"user-flow/internal/app/models"
	"user-flow/pkg/pagination"

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

func (r *Repository) Get(filter filter.GetFilter) (*pagination.Page[models.User], error) {
	base := r.db.Model(&models.User{})
	base = applyFilters(base, filter)

	var total int64
	if err := base.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("get total of users %w", err)
	}

	page := filter.Page
	if page < 1 {
		page = 1
	}

	perPage := filter.PerPage
	if filter.PerPage == 0 {
		perPage = pagination.DefaultPageSize
	}

	if page > pagination.MaxPageSize {
		page = pagination.MaxPageSize
	}

	offset := (page - 1) * perPage
	if offset < 0 {
		offset = 0
	}

	var items []models.User
	q := base

	if err := q.Order("id").
		Preload("Profile").
		Limit(perPage).
		Offset(offset).
		Find(&items).Error; err != nil {
		return nil, fmt.Errorf("failed to get users %w", err)
	}

	totalPages := int((total + int64(perPage) - 1) / int64(perPage))

	return &pagination.Page[models.User]{
		TotalItems: total,
		TotalPages: totalPages,
		PerPage:    perPage,
		Page:       page,
		Data:       items,
	}, nil
}

func applyFilters(db *gorm.DB, filter filter.GetFilter) *gorm.DB {
	db = db.Joins("JOIN user_profiles ON user_profiles.user_id = users.id")

	if filter.Keyword != "" {
		db = db.Where("user_profiles.first_name ILIKE ? OR user_profiles.last_name ILIKE ? OR users.email", filter.Keyword, filter.Keyword, filter.Keyword)
	}

	return db
}
