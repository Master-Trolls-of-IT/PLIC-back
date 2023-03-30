package services

import (
	ports "gaia-api/domain/interfaces"
)

type auth_service struct {
	db_access ports.User_interface
}

func New(db_access ports.User_interface) *auth_service {
	return &auth_service{
		db_access: db_access,
	}
}
