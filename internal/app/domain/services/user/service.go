package user

import (
	"errors"
	filter "user-flow/internal/app/domain/core/filter/user"
	repository "user-flow/internal/app/domain/repositories/user"
	"user-flow/internal/app/models"
	"user-flow/pkg/pagination"
)

var ErrNotFound = errors.New("user not found")

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Show(id int64) (models.User, error) {
	user, err := s.repo.Show(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return models.User{}, ErrNotFound
		}

		return models.User{}, err
	}

	return user, nil
}

func (s *Service) Get(filter filter.GetFilter) (*pagination.Page[models.User], error) {
	return s.repo.Get(filter)
}
