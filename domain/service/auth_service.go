package service

import (
	ports "gaia-api/domain/port"
)

type AuthService struct {
	UserRepo *ports.UserInterface
}

func NewAuthService(dbAccess ports.UserInterface) *AuthService {
	return &AuthService{
		UserRepo: &dbAccess,
	}
}
