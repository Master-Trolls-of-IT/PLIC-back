package services

import (
	ports "gaia-api/domain/interfaces"
)

type Auth_service struct {
	User_repo *ports.User_interface
}

func NewAuthService(db_access ports.User_interface) *Auth_service {
	return &Auth_service{
		User_repo: &db_access,
	}
}
