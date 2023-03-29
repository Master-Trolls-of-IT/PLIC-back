package repositories

import (
	"gaia-api/domain/entities"
)

/*type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}*/

type Error struct {
	Description string `json:"description"`
}

type User_repo struct {
	data *Database
}

func (user_repo *User_repo) getUser(query string, args ...interface{}) (entities.User, error) {
	stmt, err := user_repo.data.DB.Prepare(query)
	if err != nil {
		return entities.User{}, err
	}
	var user entities.User
	err = stmt.QueryRow(args...).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Birthdate, &user.Weight, &user.Height, &user.Gender, &user.Sportiveness, &user.BasalMetabolism)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (user_repo *User_repo) GetUserByEmail(email string) (entities.User, error) {
	return user_repo.getUser("SELECT * FROM USER WHERE Email = $1", email)
}

func (user_repo *User_repo) GetUserByUsername(username string) (entities.User, error) {
	return user_repo.getUser("SELECT * FROM USER WHERE Username = $1", username)
}


func (user_repo *User_repo) Login(login_info *entities.Login_info) {
	
}
func (user_repo *User_repo) Register(user_info *entities.User) {

}


/*func (user_auth *User_auth) Login(c *gin.Context) (entities.User, error) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		return entities.User{}, err
	}
	//Query the DB to look for the user  by username or email
	// if Query returns a row ,
	c.IndentedJSON(http.StatusFound, &Error{Description: "null"})
}*/
