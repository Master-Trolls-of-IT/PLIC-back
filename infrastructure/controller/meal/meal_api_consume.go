package meal

import (
	response "gaia-api/domain/entity/responses/meal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConsumeController struct {
	meal *Meal
}

func NewConsumeController(meal *Meal) *ConsumeController {
	consumeController := &ConsumeController{meal: meal}
	consumeController.Start()
	return consumeController
}

func (consumeController *ConsumeController) Start() {
	consumeController.meal.GinEngine.POST("/meal/consumed", consumeController.consumeMeal)
}

func (consumeController *ConsumeController) consumeMeal(context *gin.Context) {
	var meal response.Meal
	if err := context.ShouldBindJSON(&meal); err != nil {
		context.JSON(http.StatusBadRequest, consumeController.meal.ReturnAPIData.Error(http.StatusBadRequest, err.Error()))
		return
	}
	var mealRepo = *consumeController.meal.OpenFoodFactsService.MealRepo
	products, err := mealRepo.ConsumeMeal(meal)
	if err != nil {
		context.JSON(http.StatusInternalServerError, consumeController.meal.ReturnAPIData.Error(http.StatusInternalServerError, err.Error()))
	} else {
		context.JSON(http.StatusOK, consumeController.meal.ReturnAPIData.MealConsumed(products))
	}
}
