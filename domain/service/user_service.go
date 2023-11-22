package service

import (
	ports "gaia-api/domain/port"
)

type UserService struct {
	UserRepo *ports.UserInterface
}

func NewUserService(dbAccess ports.UserInterface) *UserService {
	return &UserService{
		UserRepo: &dbAccess,
	}
}
