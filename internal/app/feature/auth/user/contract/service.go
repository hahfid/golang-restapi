package contract

import (
	"context"

	"golang-restapi/internal/app/feature/auth/user/dto"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (dto.UserResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
}
