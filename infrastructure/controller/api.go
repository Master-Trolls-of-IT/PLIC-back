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
	UserService    *service.UserService
	ProductService *service.ProductService
	MealService    *service.MealService
	RecipeService  *service.RecipeService
}

func NewServer(
	userService *service.UserService,
	productService *service.ProductService,
	mealService *service.MealService,
	recipeService *service.RecipeService,
) *Server {
	return &Server{UserService: userService, ProductService: productService, MealService: mealService, RecipeService: recipeService}
}

func (server *Server) Start() {
	ginEngine := gin.Default()

	connexion.NewConnexionController(ginEngine)
	consumedProduct.NewConsumedProductController(ginEngine, server.UserService, server.ProductService)
	meal.NewMealController(ginEngine, server.UserService, server.MealService)
	recipe.NewRecipeController(ginEngine, server.UserService, server.RecipeService)
	product.NewProductController(ginEngine, server.ProductService)
	token.NewTokenController(ginEngine)
	user.NewUserController(ginEngine, server.UserService)

	err := ginEngine.Run()
	if err != nil {
		return
	}
}
