package port

import (
	"gaia-api/domain/entity/request"
	"gaia-api/domain/entity/shared"
)

type UserInterface interface {
	GetUserById(id int) (shared.User, error)
	GetUserByEmail(email string) (shared.User, error)
	GetUserByUsername(username string) (shared.User, error)
	CheckLogin(login *request.Login) (bool, error)
	Register(userInfo *shared.User) (bool, error)
	UpdateUserById(id int, newUser *shared.User) (shared.User, error)
	DeleteUser(userId int) (bool, error)
}
