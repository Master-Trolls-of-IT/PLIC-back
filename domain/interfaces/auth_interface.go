package ports

import "gaia-api/domain/entities"

type User_auth interface {
	Login(id int) (entities.User, error)
	Register(user entities.User) (int, error)
}
