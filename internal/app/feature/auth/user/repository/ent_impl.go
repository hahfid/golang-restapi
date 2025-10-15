package repository

import (
	"golang-restapi/internal/app/feature/auth/user/domain"
	"golang-restapi/internal/app/feature/auth/user/query"
)

type gormRepository struct {
	query *query.UserQuery
}

func NewGormRepository(q *query.UserQuery) UserRepository {
	return &gormRepository{query: q}
}

func (r *gormRepository) Create(user *domain.User) error {
	return r.query.Create(user)
}

func (r *gormRepository) FindByUsername(username string) (domain.User, error) {
	return r.query.FindByUsername(username)
}
