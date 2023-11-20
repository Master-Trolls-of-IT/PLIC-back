package meal

import (
	interfaces "gaia-api/application/interface"
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type Meal struct {
	GinEngine            *gin.Engine
	AuthService          *service.AuthService
	ReturnAPIData        *interfaces.ReturnAPIData
	OpenFoodFactsService *service.OpenFoodFactsService
}

func NewMealController(ginEngine *gin.Engine, authService *service.AuthService, returnAPIData *interfaces.ReturnAPIData, openFoodFactsService *service.OpenFoodFactsService) *Meal {
	meal := &Meal{GinEngine: ginEngine, AuthService: authService, ReturnAPIData: returnAPIData, OpenFoodFactsService: openFoodFactsService}
	meal.Start()
	return meal
}

func (meal *Meal) Start() {
	NewGetController(meal)
	NewAddController(meal)
	NewDeleteController(meal)
	NewConsumeController(meal)
}
