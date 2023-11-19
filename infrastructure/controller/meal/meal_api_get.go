package meal

import (
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
		context.JSON(http.StatusInternalServerError, getController.meal.ReturnAPIData.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	meals, err := mealRepo.GetMeals(user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, getController.meal.ReturnAPIData.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	context.JSON(http.StatusAccepted, getController.meal.ReturnAPIData.GetMealsSuccess(meals))

}
