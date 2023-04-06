package repositories

import (
	"gaia-api/domain/entities"
)

type Error struct {
	Description string `json:"description"`
}

type User_repo struct {
	data *Database
}

func NewUserRepository(db *Database) *User_repo {
	return &User_repo{data: db}
}
func (user_repo *User_repo) getUser(query string, args ...interface{}) (entities.User, error) {
	stmt, err := user_repo.data.DB.Prepare(query)
	if err != nil {
		return entities.User{}, err
	}
	var user entities.User
	err = stmt.QueryRow(args...).Scan(&user.Id, &user.Username,&user.Password, &user.Email, &user.Pseudo,&user.Rights, &user.Birthdate, &user.Weight, &user.Height, &user.Gender, &user.Sportiveness, &user.BasalMetabolism)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (user_repo *User_repo) GetUserByEmail(email string) (entities.User, error) {
	return user_repo.getUser("SELECT * FROM users WHERE email = $1", email)
}

func (user_repo *User_repo) GetUserByUsername(username string) (entities.User, error) {
	return user_repo.getUser("SELECT * FROM users WHERE username = $1", username)
}

func (user_repo *User_repo) GetUserById(id int) (entities.User, error) {
	return user_repo.getUser("SELECT * FROM users WHERE id = $1", id)
}

func (user_repo *User_repo) CheckLogin(login_info *entities.Login_info) (bool, error) {
	user, err := user_repo.getUser("SELECT * FROM users WHERE username=$1 OR email=$2", login_info.Username, login_info.Email)
	if err != nil {
		return false, err
	}
	return user.Password == login_info.Password, nil
}
func (user_repo *User_repo) Register(user_info *entities.User) (bool, error) {
	var db = user_repo.data.DB
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username=$1 OR email=$2", user_info.Username, user_info.Email).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return false, nil
	}
	_, err = db.Exec("INSERT INTO users (email, rights,username, password, birthdate, weight, height, gender, sportiveness, basalmetabolism) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		user_info.Email,user_info.. user_info.Username, user_info.Password, user_info.Birthdate, user_info.Weight, user_info.Height, user_info.Gender, user_info.Sportiveness, user_info.BasalMetabolism)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (user_repo *User_repo) Login(login_info *entities.Login_info) {}
