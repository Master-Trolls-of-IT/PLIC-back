package interfaces

import (
	"fmt"
	"gaia-api/domain/entity"
)

type IReturnAPIData interface {
	Error(errorStatus int, errorMessage string)
	LoginSuccess(user entity.User)
	RegisterSucces(user entity.User)
	GetToken(token string)
	CheckToken(isTokenValid bool)
	ProductFound(product entity.Product)
	ProductNotAvailable(barcode string)
	Ping()
}

type ReturnAPIData struct {
}

func NewReturnAPIData() *ReturnAPIData {
	return &ReturnAPIData{}
}

type JSONObject map[string]any
type JSONList []map[string]any

func (returnAPIData *ReturnAPIData) Error(errorStatus int, errorMessage string) JSONObject {
	return JSONObject{
		"status":  errorStatus,
		"message": errorMessage,
		"data":    JSONObject{},
	}
}

func (returnAPIData *ReturnAPIData) LoginSuccess(user entity.User) JSONObject {
	return JSONObject{
		"status":  202,
		"message": "Connecté avec succès",
		"data": JSONObject{
			"Id":              user.Id,
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
			"AvatarId":        user.AvatarId,
		},
	}
}

func (ReturnAPIData *ReturnAPIData) ValidPassword(message string) JSONObject {
	return JSONObject{
		"status":  202,
		"message": message,
	}
}

func (returnAPIData *ReturnAPIData) RegisterSuccess(user entity.User) JSONObject {
	return JSONObject{
		"status":  200,
		"message": "Enregistré avec succès",
		"data": JSONObject{
			"Id":              user.Id,
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
			"AvatarId":        user.AvatarId,
		},
	}
}

func (returnAPIData *ReturnAPIData) UserUpdateSuccess(user entity.User) JSONObject {
	return JSONObject{
		"status":  200,
		"message": "Modifié avec succès",
		"data": JSONObject{
			"Id":              user.Id,
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
			"AvatarId":        user.AvatarId,
		},
	}
}

func (returnAPIData *ReturnAPIData) GetToken(token string) JSONObject {
	return JSONObject{
		"status":  200,
		"message": "Token généré avec succès",
		"data": JSONObject{
			"token": token,
		},
	}
}

func (returnAPIData *ReturnAPIData) CheckToken(isTokenValid bool) JSONObject {
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

func (returnAPIData *ReturnAPIData) ProductFound(product entity.Product) JSONObject {
	return JSONObject{
		"status":  200,
		"message": "Les informations du produit ont été trouvées avec succès",
		"data":    product,
	}
}

func (returnAPIData *ReturnAPIData) ProductAddedToConsumed(product entity.Product) JSONObject {
	return JSONObject{
		"status":  200,
		"message": "Le produit a été ajouté avec succès",
		"data":    product,
	}
}

func (returnAPIData *ReturnAPIData) ProductNotAvailable(barcode string) JSONObject {
	return JSONObject{
		"status":  404,
		"message": fmt.Sprintf("Le produit de barcode: %s n'est pas disponible", barcode),
		"data":    "",
	}
}

func (returnAPIData *ReturnAPIData) Ping() JSONObject {
	return JSONObject{
		"status":  200,
		"message": "Pong",
		"data":    JSONObject{},
	}
}

func (returnAPIData *ReturnAPIData) GetConsumedProductsSuccess(consumedProducts []entity.ConsumedProduct) any {
	return JSONObject{
		"status":  200,
		"message": "Les produits consommés ont été récupérés avec succès",
		"data":    consumedProducts,
	}
}

func (returnAPIData *ReturnAPIData) ProductDeletedFromConsumed(consumedProducts []entity.ConsumedProduct) any {
	return JSONObject{
		"status":  200,
		"message": "Le produit a été supprimé avec succès",
		"data":    consumedProducts,
	}
}

func (returnAPIData *ReturnAPIData) DeletedProduct(status int, s string) any {
	return JSONObject{
		"status":  status,
		"message": s,
		"data":    JSONObject{},
	}
}
