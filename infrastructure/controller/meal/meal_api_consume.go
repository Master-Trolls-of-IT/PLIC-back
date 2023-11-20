package meal

import (
	"gaia-api/application/returnAPI"
	"gaia-api/domain/entity/response"
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
		returnAPI.Error(context, http.StatusBadRequest)
		return
	}
	var mealRepo = *consumeController.meal.OpenFoodFactsService.MealRepo
	products, err := mealRepo.ConsumeMeal(meal)
	if err != nil {
		returnAPI.Error(context, http.StatusInternalServerError)
	} else {
		returnAPI.Success(context, http.StatusOK, products)
	}
}
