package meal

import (
	"gaia-api/domain/service"
	"github.com/gin-gonic/gin"
)

type Meal struct {
	GinEngine   *gin.Engine
	UserService *service.UserService
	MealService *service.MealService
}

func NewMealController(ginEngine *gin.Engine, userService *service.UserService, mealService *service.MealService) *Meal {
	meal := &Meal{GinEngine: ginEngine, UserService: userService, MealService: mealService}
	meal.Start()
	return meal
}

func (meal *Meal) Start() {
	NewGetController(meal)
	NewAddController(meal)
	NewDeleteController(meal)
	NewConsumeController(meal)
}
