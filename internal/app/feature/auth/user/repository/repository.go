package repository

import (
	"golang-restapi/internal/app/feature/auth/user/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByUsername(username string) (domain.User, error)
}
