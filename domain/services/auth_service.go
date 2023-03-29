package services

import (
	entities "gaia-api/domain/entities"
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

func (service *auth_service) Login(id int) (entities.User, error) {
}

func (service *auth_service) Register(user entities.User) (int, error) {
}
