package meal

import (
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type Meal struct {
	GinEngine            *gin.Engine
	AuthService          *service.AuthService
	OpenFoodFactsService *service.OpenFoodFactsService
}

func NewMealController(ginEngine *gin.Engine, authService *service.AuthService, openFoodFactsService *service.OpenFoodFactsService) *Meal {
	meal := &Meal{GinEngine: ginEngine, AuthService: authService, OpenFoodFactsService: openFoodFactsService}
	meal.Start()
	return meal
}

func (meal *Meal) Start() {
	NewGetController(meal)
	NewAddController(meal)
	NewDeleteController(meal)
	NewConsumeController(meal)
}
