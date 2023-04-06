package ports

import "gaia-api/domain/entities"

type User_interface interface {
	GetUserById(id int) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	GetUserByUsername(username string) (entities.User, error)

	CheckLogin(login_info *entities.Login_info) (bool, error)
	Register(user_info *entities.User) (bool, error)
}
