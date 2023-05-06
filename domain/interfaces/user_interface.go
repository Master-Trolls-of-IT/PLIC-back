package ports

import "gaia-api/domain/entities"

type UserInterface interface {
	GetUserById(id int) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	GetUserByUsername(username string) (entities.User, error)

	CheckLogin(loginInfo *entities.Login_info) (bool, error)
	Register(userInfo *entities.User) (bool, error)
}
