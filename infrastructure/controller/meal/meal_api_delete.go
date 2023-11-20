package meal

import (
	"gaia-api/application/returnAPI"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteController struct {
	meal *Meal
}

func NewDeleteController(meal *Meal) *DeleteController {
	deleteController := &DeleteController{meal: meal}
	deleteController.Start()
	return deleteController
}

func (deleteController *DeleteController) Start() {
	deleteController.meal.GinEngine.DELETE("/meal/:id", deleteController.deleteMeal)
}

func (deleteController *DeleteController) deleteMeal(context *gin.Context) {
	mealID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		returnAPI.Error(context, http.StatusBadRequest)
		return
	}
	var mealRepo = *deleteController.meal.OpenFoodFactsService.MealRepo

	err = mealRepo.DeleteMeal(mealID)
	if err != nil {
		returnAPI.Error(context, http.StatusInternalServerError)
		return
	}
	returnAPI.Success(context, http.StatusOK, nil)
}
