package mapper

import (
	"golang-restapi/internal/app/feature/auth/user/domain"
	"golang-restapi/internal/app/feature/auth/user/dto"
)

func ToUserResponse(user domain.User) dto.UserResponse {
	return dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}
