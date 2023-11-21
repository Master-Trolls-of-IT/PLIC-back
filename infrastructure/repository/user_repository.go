package repository

import (
	"gaia-api/domain/entity/request"
	"gaia-api/domain/entity/shared"
)

type UserRepo struct {
	data *Database
}

func NewUserRepository(data *Database) *UserRepo {
	return &UserRepo{data: data}
}

func (userRepo *UserRepo) getUser(query string, args ...interface{}) (shared.User, error) {
	stmt, err := userRepo.data.DB.Preparex(query)
	if err != nil {
		return shared.User{}, err
	}

	var user shared.User
	err = stmt.Get(&user, args...)
	if err != nil {
		return shared.User{}, err
	}

	return user, nil
}

func (userRepo *UserRepo) GetUserByEmail(email string) (shared.User, error) {
	return userRepo.getUser("SELECT * FROM users WHERE email = $1", email)
}

func (userRepo *UserRepo) GetUserByUsername(username string) (shared.User, error) {
	return userRepo.getUser("SELECT * FROM users WHERE username = $1", username)
}

func (userRepo *UserRepo) GetUserById(id int) (shared.User, error) {
	return userRepo.getUser("SELECT * FROM users WHERE id = $1", id)
}

func (userRepo *UserRepo) CheckLogin(login *request.Login) (bool, error) {
	user, err := userRepo.getUser("SELECT * FROM users WHERE username=$1 OR email=$2", login.Username, login.Email)
	if err != nil {
		return false, err
	}
	return user.Password == login.Password, nil
}

func (userRepo *UserRepo) Register(userInfo *shared.User) (bool, error) {
	var database = userRepo.data.DB

	var count int
	var query = "SELECT COUNT(*) FROM users WHERE username=$1 OR email=$2"
	err := database.Get(&count, query, userInfo.Username, userInfo.Email)
	if err != nil || count > 0 {
		return false, err
	}

	userInsertQuery := `
        INSERT INTO users (email, pseudo, rights, username, password, birthdate, weight, height, gender, sportiveness, basalmetabolism, avatar_id) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id;
    `
	err = database.QueryRow(userInsertQuery, userInfo.Email, userInfo.Pseudo, userInfo.Rights, userInfo.Username,
		userInfo.Password, userInfo.Birthdate, userInfo.Weight, userInfo.Height, userInfo.Gender, userInfo.Sportiveness,
		userInfo.BasalMetabolism, userInfo.AvatarId).Scan(&userInfo.Id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (userRepo *UserRepo) UpdateUserById(id int, newUser *shared.User) (shared.User, error) {
	var db = userRepo.data.DB
	stmt, err := db.Preparex(`
		UPDATE users SET email = $1, pseudo = $2, birthdate = $3, weight = $4, height = $5, gender = $6,
		sportiveness = $7, basalmetabolism = $8, avatar_id = $9 WHERE id = $10
	`)

	if err != nil {
		return shared.User{}, err
	}
	_, err = stmt.Exec(newUser.Email,
		&newUser.Pseudo, &newUser.Birthdate, &newUser.Weight, &newUser.Height, &newUser.Gender, &newUser.Sportiveness,
		&newUser.BasalMetabolism, &newUser.AvatarId, &id)

	if err != nil {
		return shared.User{}, err
	}
	return userRepo.GetUserById(id)
}

func (userRepo *UserRepo) DeleteUser(userId int) (bool, error) {
	var database = userRepo.data.DB
	_, err := database.Exec("DELETE FROM users WHERE id = $1", userId)
	if err != nil {
		return false, err
	}
	return true, nil
}
