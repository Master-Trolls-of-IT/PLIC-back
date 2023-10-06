package repository

import (
	"gaia-api/domain/entity"
)

type Error struct {
	Description string `json:"description"`
}

type UserRepo struct {
	data *Database
}

func NewUserRepository(db *Database) *UserRepo {
	return &UserRepo{data: db}
}

func (userRepo *UserRepo) getUser(query string, args ...interface{}) (entity.User, error) {
	stmt, err := userRepo.data.DB.Prepare(query)
	if err != nil {
		return entity.User{}, err
	}
	var user entity.User
	err = stmt.QueryRow(args...).Scan(&user.Id, &user.Rights, &user.Email, &user.Username, &user.Birthdate, &user.Weight,
		&user.Height, &user.Gender, &user.Sportiveness, &user.BasalMetabolism, &user.Password, &user.Pseudo, &user.AvatarId)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (userRepo *UserRepo) GetUserByEmail(email string) (entity.User, error) {
	return userRepo.getUser("SELECT * FROM users WHERE email = $1", email)
}

func (userRepo *UserRepo) GetUserByUsername(username string) (entity.User, error) {
	return userRepo.getUser("SELECT * FROM users WHERE username = $1", username)
}

func (userRepo *UserRepo) GetUserById(id int) (entity.User, error) {
	return userRepo.getUser("SELECT * FROM users WHERE id = $1", id)
}

func (userRepo *UserRepo) CheckLogin(loginInfo *entity.Login_info) (bool, error) {
	user, err := userRepo.getUser("SELECT * FROM users WHERE username=$1 OR email=$2", loginInfo.Username, loginInfo.Email)
	if err != nil {
		return false, err
	}
	return user.Password == loginInfo.Password, nil
}
func (userRepo *UserRepo) Register(userInfo *entity.User) (bool, error) {
	var db = userRepo.data.DB
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username=$1 OR email=$2", userInfo.Username,
		userInfo.Email).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return false, nil
	}
	_, err = db.Exec("INSERT INTO users (email, pseudo, rights,username, password, birthdate, weight, height, "+
		"gender, sportiveness, basalmetabolism) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		userInfo.Email, userInfo.Pseudo, userInfo.Rights, userInfo.Username, userInfo.Password, userInfo.Birthdate,
		userInfo.Weight, userInfo.Height, userInfo.Gender, userInfo.Sportiveness, userInfo.BasalMetabolism)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (userRepo *UserRepo) UpdateUserById(id int, newUser *entity.User) (entity.User, error) {
	var db = userRepo.data.DB
	stmt, err := db.Prepare("UPDATE users SET email = $1, pseudo = $2, birthdate = $3, weight = $4, " +
		"height = $5, gender = $6, sportiveness = $7, basalmetabolism = $8, avatar_id = $9 WHERE id = $10")

	if err != nil {
		return entity.User{}, err
	}
	_, err = stmt.Exec(newUser.Email,
		&newUser.Pseudo, &newUser.Birthdate, &newUser.Weight, &newUser.Height, &newUser.Gender, &newUser.Sportiveness,
		&newUser.BasalMetabolism, &newUser.AvatarId, &id)

	if err != nil {
		return entity.User{}, err
	}
	return userRepo.GetUserById(id)
}

func (userRepo *UserRepo) Login(loginInfo *entity.Login_info) {}
