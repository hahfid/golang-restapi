package query

import (
	"golang-restapi/internal/app/feature/auth/user/domain"

	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func NewUserQuery(db *gorm.DB) *UserQuery {
	return &UserQuery{db: db}
}

func (q *UserQuery) FindByUsername(username string) (domain.User, error) {
	var user domain.User
	err := q.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (q *UserQuery) Create(user *domain.User) error {
	return q.db.Create(user).Error
}
