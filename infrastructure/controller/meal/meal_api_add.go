package meal

import (
	"gaia-api/application/returnAPI"
	"gaia-api/domain/entity/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddController struct {
	meal *Meal
}

func NewAddController(meal *Meal) *AddController {
	addController := &AddController{meal: meal}
	addController.Start()
	return addController
}

func (addController *AddController) Start() {
	addController.meal.GinEngine.POST("/meal", addController.addMeal)
}

func (addController *AddController) addMeal(context *gin.Context) {
	var meal request.Meal
	if err := context.ShouldBindJSON(&meal); err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
		return
	}

	var mealRepo = *addController.meal.OpenFoodFactsService.MealRepo

	responseMeal, err := mealRepo.SaveMeal(meal)
	if err != nil {
		returnAPI.Error(context, http.StatusInternalServerError)
	} else {
		returnAPI.Success(context, http.StatusCreated, *responseMeal)
	}
}
