package services

import (
	ports "gaia-api/domain/interfaces"
)

type AuthService struct {
	UserRepo *ports.UserInterface
}

func NewAuthService(dbAccess ports.UserInterface) *AuthService {
	return &AuthService{
		UserRepo: &dbAccess,
	}
}
