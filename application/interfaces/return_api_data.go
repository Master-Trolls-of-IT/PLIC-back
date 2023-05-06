package interfaces

import (
	"gaia-api/domain/entities"
)

type IReturnAPIData interface {
	Error(errorStatus int, errorMessage string)
	LoginSuccess(user entities.User)
	RegisterSucces(user entities.User)
	GetToken(token string)
	CheckToken(isTokenValid bool)
}

type ReturnAPIData struct {
}

func NewReturnAPIData() *ReturnAPIData {
	return &ReturnAPIData{}
}

type JSONObject map[string]any
type JSONList []map[string]any

func (ReturnAPIData *ReturnAPIData) Error(errorStatus int, errorMessage string) JSONObject {
	return JSONObject{
		"status":  errorStatus,
		"message": errorMessage,
		"data":    JSONObject{},
	}
}

func (ReturnAPIData *ReturnAPIData) LoginSuccess(user entities.User) JSONObject {
	return JSONObject{
		"status":  202,
		"message": "Connecté avec succès",
		"data": JSONObject{
			"Email":           user.Email,
			"Username":        user.Username,
			"Birthdate":       user.Birthdate,
			"Weight":          user.Weight,
			"Height":          user.Height,
			"Gender":          user.Gender,
			"Pseudo":          user.Pseudo,
			"Rights":          user.Rights,
			"Sportiveness":    user.Sportiveness,
			"BasalMetabolism": user.BasalMetabolism,
		},
	}
}

func (ReturnAPIData *ReturnAPIData) RegisterSuccess(user entities.User) JSONObject {
	return JSONObject{
		"status":  200,
		"message": "Enregistré avec succès",
		"data": JSONObject{
			"Email":           user.Email,
			"Username":        user.Username,
			"Birthdate":       user.Birthdate,
			"Weight":          user.Weight,
			"Height":          user.Height,
			"Gender":          user.Gender,
			"Pseudo":          user.Pseudo,
			"Rights":          user.Rights,
			"Sportiveness":    user.Sportiveness,
			"BasalMetabolism": user.BasalMetabolism,
		},
	}
}

func (ReturnAPIData *ReturnAPIData) GetToken(token string) JSONObject {
	return JSONObject{
		"status":  200,
		"message": "Token généré avec succès",
		"data": JSONObject{
			"token": token,
		},
	}
}

func (ReturnAPIData *ReturnAPIData) CheckToken(isTokenValid bool) JSONObject {
	getMessage := func(messageBool bool) string {
		if messageBool {
			return "Le token est valide"
		}
		return "Le token n'est pas valide"
	}

	return JSONObject{
		"status":  200,
		"message": getMessage(isTokenValid),
		"data": JSONObject{
			"isTokenValid": isTokenValid,
		},
	}
}
