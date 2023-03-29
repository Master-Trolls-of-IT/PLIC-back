package ports

import "gaia-api/domain/entities"



type User_interface interface {
	GetUserById(id int) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	GetUserByUsername(username string) (entities.User, error)

	Login(login_info *entities.Login_info)
	Register(user_info *entities.User)
}
