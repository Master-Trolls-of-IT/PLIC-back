package port

import (
	"gaia-api/domain/entity"
)

type UserInterface interface {
	GetUserById(id int) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	CheckLogin(loginInfo *entity.Login_info) (bool, error)
	Register(userInfo *entity.User) (bool, error)
	UpdateUserById(id int, newUser *entity.User) (entity.User, error)
	DeleteUser(userId int) (bool, error)
}
