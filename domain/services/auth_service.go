package services

import (
	entities "gaia-api/domain/entities"
	ports "gaia-api/domain/interfaces"
)

type auth_service struct {
	db_access ports.User_auth
}

func New(db_access ports.User_auth) *auth_service {
	return &auth_service{
		db_access: db_access,
	}
}

func (service *auth_service) Login(id int) (entities.User, error) {
	user, err := service.Login(id)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (service *auth_service) Register(user entities.User) (int, error) {
	user_id, err := service.Register(user)
	if err != nil {
		return 0, err
	}
	return user_id, nil
}
