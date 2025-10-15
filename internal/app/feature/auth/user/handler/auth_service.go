package handler

import (
	"context"
	"errors"

	"golang-restapi/internal/app/feature/auth/user/contract"
	"golang-restapi/internal/app/feature/auth/user/domain"
	"golang-restapi/internal/app/feature/auth/user/dto"
	"golang-restapi/internal/app/feature/auth/user/mapper"
	"golang-restapi/internal/app/feature/auth/user/repository"
	"golang-restapi/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo   repository.UserRepository
	tokens *jwt.TokenManager
}

func NewAuthService(repo repository.UserRepository, tokens *jwt.TokenManager) contract.AuthService {
	return &authService{repo: repo, tokens: tokens}
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequest) (dto.UserResponse, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserResponse{}, err
	}

	user := domain.User{
		Username: req.Username,
		Password: string(hashed),
	}

	if err := s.repo.Create(&user); err != nil {
		return dto.UserResponse{}, err
	}

	return mapper.ToUserResponse(user), nil
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}

	token, err := s.tokens.GenerateToken(user.ID)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{Token: token}, nil
}
