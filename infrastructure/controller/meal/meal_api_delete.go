package meal

import (
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
		context.JSON(http.StatusBadRequest, deleteController.meal.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
		return
	}
	var mealRepo = *deleteController.meal.OpenFoodFactsService.MealRepo

	err = mealRepo.DeleteMeal(mealID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, deleteController.meal.ReturnAPIData.Error(http.StatusInternalServerError, err.Error()))
		return
	}
	context.JSON(http.StatusOK, deleteController.meal.ReturnAPIData.MealDeleted())
}
