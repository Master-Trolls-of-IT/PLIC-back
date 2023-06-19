package controller

import (
	"database/sql"
	"fmt"
	"gaia-api/application/interface"
	"gaia-api/domain/entity"
	"gaia-api/domain/service"
	"gaia-api/infrastructure/error/openFoodFacts_api_error"
	"net/http"

	_ "github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

type Server struct {
	authService          *service.AuthService
	openFoodFactsService *service.OpenFoodFactsService
	OpenFoodFactsAPI     *OpenFoodFactsAPI

	returnAPIData *interfaces.ReturnAPIData
	// TODO: Store the logs ?
	//logger        *service.LoggerService
}

func NewServer(authService *service.AuthService, returnAPIData *interfaces.ReturnAPIData, openFoodFactsService *service.OpenFoodFactsService, OpenFoodFactsAPI *OpenFoodFactsAPI) *Server {
	return &Server{authService: authService, returnAPIData: returnAPIData, openFoodFactsService: openFoodFactsService, OpenFoodFactsAPI: OpenFoodFactsAPI}
}

func (server *Server) Start() {
	ginEngine := gin.Default()

	ginEngine.GET("/", server.welcome)
	ginEngine.GET("/ping", server.ping)

	ginEngine.POST("/logs", server.printLogs)

	ginEngine.POST("/login", server.login)
	ginEngine.POST("/register", server.register)
	ginEngine.PUT("/users/:id", server.updateProfile)
	ginEngine.DELETE("/users/:id", server.deleteAccount)

	ginEngine.GET("/refresh_token/:password", server.getRefreshToken)
	ginEngine.GET("/access_token/:password/:refreshtoken", server.getAccessToken)
	ginEngine.GET("/access_token/check/:token", server.checkAccessToken)
	ginEngine.GET("/refresh_token/check/:token", server.checkRefreshToken)

	ginEngine.GET("/product/:barcode", server.mapAndSaveAndGetProduct)

	err := ginEngine.Run()
	if err != nil {
		return
	}
}

func (server *Server) mapAndSaveAndGetProduct(context *gin.Context) {
	var barcode = context.Param("barcode")
	var productRepo = *server.openFoodFactsService.ProductRepo
	product, dbError := productRepo.GetProductByBarCode(barcode)
	fmt.Print(product)

	if dbError != nil && dbError != sql.ErrNoRows {
		context.JSON(http.StatusInternalServerError, server.returnAPIData.Error(http.StatusInternalServerError, dbError.Error()))

	} else if product == (entity.Product{}) {

		openFoodFactAPI := server.OpenFoodFactsAPI
		mappedProduct, err := openFoodFactAPI.retrieveAndMapProduct(barcode)

		if _, productNotFound := err.(openFoodFacts_api_error.ProductNotFoundError); productNotFound {
			context.JSON(http.StatusInternalServerError, server.returnAPIData.ProductNotAvailable(barcode))

		} else {
			productSaved, err := productRepo.SaveProduct(mappedProduct, barcode)
			if productSaved {
				context.JSON(http.StatusOK, server.returnAPIData.ProductFound(mappedProduct))
			} else {
				context.JSON(http.StatusInternalServerError, server.returnAPIData.Error(http.StatusInternalServerError, err.Error()))
			}
		}

	} else {
		context.JSON(http.StatusOK, server.returnAPIData.ProductFound(product))
	}
}

func (server *Server) printLogs(context *gin.Context) {
	var logs []entity.UserLogs
	var color string
	if err := context.BindJSON(&logs); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	for _, log := range logs {
		switch log.Level {
		case "ERROR":
			color = "\033[31m" // red
		case "WARNING":
			color = "\033[33m" // yellow
		case "INFO":
			color = "\033[32m" // green
		default:
			color = "\033[0m" // reset color
		}
		fmt.Printf("\n%s[%s] {%s}:\nDescription: %s\nDetails: %s\nSource: %s\n", color, log.Date, log.Level, log.Message, log.Details, log.Source)
	}
}

// Function that checks if the access token is valid, it takes an access token as parameter and returns a JSON with this structure: {"valid": true}
func (server *Server) checkAccessToken(context *gin.Context) {
	isTokenValid, err := verifyAccessToken(context.Param("token"))
	if err != nil {
		context.JSON(http.StatusBadRequest, server.returnAPIData.Error(http.StatusBadRequest, err.Error()))
	} else {
		context.JSON(http.StatusOK, server.returnAPIData.CheckToken(isTokenValid))
	}
}

// Function that checks if the refresh token is valid, it takes a refresh token as parameter and returns a JSON with this structure: {"valid": true}
func (server *Server) checkRefreshToken(context *gin.Context) {
	isTokenValid, err := verifyRefreshToken(context.Param("token"))
	if err != nil {
		context.JSON(http.StatusBadRequest, server.returnAPIData.Error(http.StatusBadRequest, err.Error()))
	} else {
		context.JSON(http.StatusOK, server.returnAPIData.CheckToken(isTokenValid))
	}
}

// Function that generates an access token, it takes a password and a refreshtoken as parameter and returns a JSON with this structure : { "token": generatedToken }
func (server *Server) getAccessToken(context *gin.Context) {
	// Use GenerateAccessToken function to generate a new access token
	accessToken, err := generateAccessToken(context.Param("password"), []byte(context.Param("refreshtoken")))
	if err != nil {
		context.JSON(http.StatusBadRequest, server.returnAPIData.Error(http.StatusBadRequest, err.Error()))
	}
	context.JSON(http.StatusOK, server.returnAPIData.GetToken(accessToken))
}

// Function that generates a refresh token, it takes a password as parameter and returns a JSON with this structure : { "token": generatedToken }
func (server *Server) getRefreshToken(context *gin.Context) {
	// Use GenerateRefreshToken function to generate a new refresh token
	refreshToken, err := generateRefreshToken([]byte(context.Param("password")))
	if err != nil {
		context.JSON(http.StatusBadRequest, server.returnAPIData.Error(http.StatusBadRequest, err.Error()))
	}
	context.JSON(http.StatusOK, server.returnAPIData.GetToken(refreshToken))
}

// Welcome function that returns a JSON with this structure : { "Title": "Gaia" }
func (server *Server) welcome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Title": "Gaia"})
}

// Ping function that returns a JSON with this structure : { "ping": "pong" }
func (server *Server) ping(context *gin.Context) {
	context.JSON(http.StatusOK, server.returnAPIData.Ping())
}

func (server *Server) login(context *gin.Context) {
	var login = entity.Login_info{}
	//binds Json Body to Entities.Login_info Class
	if err := context.ShouldBindJSON(&login); err != nil {
		context.JSON(http.StatusBadRequest, server.returnAPIData.Error(http.StatusBadRequest, err.Error()))
	}
	var userRepo = *server.authService.UserRepo
	loggedIn, err := userRepo.CheckLogin(&login)
	if err != nil {
		context.JSON(http.StatusUnauthorized, server.returnAPIData.Error(http.StatusUnauthorized, "Informations de connexion non valides"))
	} else if loggedIn {
		var user entity.User
		if login.Email == "" {
			user, _ = userRepo.GetUserByUsername(login.Username)
		} else {
			user, _ = userRepo.GetUserByEmail(login.Email)
		}

		//a function that generates a token using JWT
		context.JSON(http.StatusAccepted, server.returnAPIData.LoginSuccess(user))
	} else {
		context.JSON(http.StatusInternalServerError, server.returnAPIData.Error(http.StatusBadRequest, err.Error()))
	}
}

func (server *Server) register(context *gin.Context) {
	var user = entity.User{}
	//binds Json Body to Entities.User Class
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, server.returnAPIData.Error(http.StatusBadRequest, err.Error()))
		return
	}

	var userRepo = *server.authService.UserRepo
	registered, err := userRepo.Register(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, server.returnAPIData.Error(http.StatusInternalServerError, err.Error()))
	} else if registered {
		context.JSON(http.StatusOK, server.returnAPIData.RegisterSuccess(user))
	} else {
		context.JSON(http.StatusConflict, server.returnAPIData.Error(http.StatusConflict, "Nom d'utilisateur ou email déjà utilisée"))
	}
}

func (server *Server) updateProfile(context *gin.Context) {
	// TODO: Update user profile
}

func (server *Server) deleteAccount(context *gin.Context) {
	// TODO: Delete account
}
