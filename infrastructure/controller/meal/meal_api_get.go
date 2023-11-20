package meal

import (
	"gaia-api/application/returnAPI"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetController struct {
	meal *Meal
}

func NewGetController(meal *Meal) *GetController {
	getController := &GetController{meal: meal}
	getController.Start()
	return getController
}

func (getController *GetController) Start() {
	getController.meal.GinEngine.GET("/meal/:email", getController.getMeals)
}

func (getController *GetController) getMeals(context *gin.Context) {
	var email = context.Param("email")
	var userRepo, mealRepo = *getController.meal.AuthService.UserRepo, *getController.meal.OpenFoodFactsService.MealRepo

	user, err := userRepo.GetUserByEmail(email)
	if err != nil {
		returnAPI.Error(context, http.StatusInternalServerError)
		return
	}

	meals, err := mealRepo.GetMeals(user.Email)
	if err != nil {
		returnAPI.Error(context, http.StatusInternalServerError)
		return
	} else if len(meals) == 0 {
		returnAPI.Success(context, http.StatusOK, nil)
	}
	returnAPI.Success(context, http.StatusOK, meals)
}
