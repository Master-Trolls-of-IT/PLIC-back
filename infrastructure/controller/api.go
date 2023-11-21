package controller

import (
	"gaia-api/domain/service"
	"gaia-api/infrastructure/controller/connexion"
	"gaia-api/infrastructure/controller/consumedProduct"
	"gaia-api/infrastructure/controller/meal"
	"gaia-api/infrastructure/controller/product"
	"gaia-api/infrastructure/controller/recipe"
	"gaia-api/infrastructure/controller/token"
	"gaia-api/infrastructure/controller/user"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
)

type Server struct {
	UserService      *service.UserService
	ProductService   *service.ProductService
	MealService      *service.MealService
	OpenFoodFactsAPI *product.OpenFoodFactsAPI
}

func NewServer(
	userService *service.UserService,
	productService *service.ProductService,
	mealService *service.MealService,
	OpenFoodFactsAPI *product.OpenFoodFactsAPI,
) *Server {
	return &Server{UserService: userService, ProductService: productService, MealService: mealService, OpenFoodFactsAPI: OpenFoodFactsAPI}
}

func (server *Server) Start() {
	ginEngine := gin.Default()

	connexion.NewConnexionController(ginEngine)
	consumedProduct.NewConsumedProductController(ginEngine, server.UserService, server.ProductService)
	meal.NewMealController(ginEngine, server.UserService, server.MealService)
	recipe.NewRecipeController(ginEngine, server.UserService)
	product.NewProductController(ginEngine, server.ProductService, server.OpenFoodFactsAPI)
	token.NewTokenController(ginEngine)
	user.NewUserController(ginEngine, server.UserService)

	err := ginEngine.Run()
	if err != nil {
		return
	}
}
