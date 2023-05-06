package repositories

import (
	"gaia-api/domain/entities"
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
func (userRepo *UserRepo) getUser(query string, args ...interface{}) (entities.User, error) {
	stmt, err := userRepo.data.DB.Prepare(query)
	if err != nil {
		return entities.User{}, err
	}
	var user entities.User
	err = stmt.QueryRow(args...).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Pseudo, &user.Rights, &user.Birthdate, &user.Weight, &user.Height, &user.Gender, &user.Sportiveness, &user.BasalMetabolism)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (userRepo *UserRepo) GetUserByEmail(email string) (entities.User, error) {
	return userRepo.getUser("SELECT * FROM users WHERE email = $1", email)
}

func (userRepo *UserRepo) GetUserByUsername(username string) (entities.User, error) {
	return userRepo.getUser("SELECT * FROM users WHERE username = $1", username)
}

func (userRepo *UserRepo) GetUserById(id int) (entities.User, error) {
	return userRepo.getUser("SELECT * FROM users WHERE id = $1", id)
}

func (userRepo *UserRepo) CheckLogin(loginInfo *entities.Login_info) (bool, error) {
	user, err := userRepo.getUser("SELECT * FROM users WHERE username=$1 OR email=$2", loginInfo.Username, loginInfo.Email)
	if err != nil {
		return false, err
	}
	return user.Password == loginInfo.Password, nil
}
func (userRepo *UserRepo) Register(userInfo *entities.User) (bool, error) {
	var db = userRepo.data.DB
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username=$1 OR email=$2", userInfo.Username, userInfo.Email).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return false, nil
	}
	_, err = db.Exec("INSERT INTO users (email, pseudo, rights,username, password, birthdate, weight, height, gender, sportiveness, basalmetabolism) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		userInfo.Email, userInfo.Pseudo, userInfo.Rights, userInfo.Username, userInfo.Password, userInfo.Birthdate, userInfo.Weight, userInfo.Height, userInfo.Gender, userInfo.Sportiveness, userInfo.BasalMetabolism)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (userRepo *UserRepo) Login(loginInfo *entities.Login_info) {}
