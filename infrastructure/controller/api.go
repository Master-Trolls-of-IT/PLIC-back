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
	AuthService          *service.AuthService
	OpenFoodFactsService *service.OpenFoodFactsService
	OpenFoodFactsAPI     *product.OpenFoodFactsAPI
}

func NewServer(authService *service.AuthService, OpenFoodFactsService *service.OpenFoodFactsService, OpenFoodFactsAPI *product.OpenFoodFactsAPI) *Server {
	return &Server{AuthService: authService, OpenFoodFactsService: OpenFoodFactsService, OpenFoodFactsAPI: OpenFoodFactsAPI}
}

func (server *Server) Start() {
	ginEngine := gin.Default()

	connexion.NewConnexionController(ginEngine)
	consumedProduct.NewConsumedProductController(ginEngine, server.AuthService, server.OpenFoodFactsService)
	meal.NewMealController(ginEngine, server.AuthService, server.OpenFoodFactsService)
	recipe.NewRecipeController(ginEngine, server.AuthService)
	product.NewProductController(ginEngine, server.OpenFoodFactsService, server.OpenFoodFactsAPI)
	token.NewTokenController(ginEngine)
	user.NewUserController(ginEngine, server.AuthService)

	err := ginEngine.Run()
	if err != nil {
		return
	}
}
