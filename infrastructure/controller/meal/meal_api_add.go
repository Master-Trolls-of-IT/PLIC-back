package meal

import (
	"gaia-api/domain/entity/requests/mealRequest"
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
	var meal mealRequest.Meal
	if err := context.ShouldBindJSON(&meal); err != nil {
		context.JSON(http.StatusBadRequest, addController.meal.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
		return
	}

	var mealRepo = *addController.meal.OpenFoodFactsService.MealRepo

	responseMeal, err := mealRepo.SaveMeal(meal)
	if err != nil {
		context.JSON(http.StatusInternalServerError, addController.meal.ReturnAPIData.Error(http.StatusInternalServerError, err.Error()))
	} else {
		context.JSON(http.StatusOK, addController.meal.ReturnAPIData.MealAdded(*responseMeal))
	}
}
