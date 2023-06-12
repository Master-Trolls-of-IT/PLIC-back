package controllers

import (
	"fmt"
	"gaia-api/application/interfaces"
	"gaia-api/domain/entities"
	"gaia-api/domain/services"
	"net/http"

	_ "github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

type Server struct {
	authService          *services.AuthService
	openFoodFactsService *services.OpenFoodFactsService

	returnAPIData *interfaces.ReturnAPIData
	// TODO: Store the logs ?
	//logger        *services.LoggerService
}

func NewServer(authService *services.AuthService, returnAPIData *interfaces.ReturnAPIData) *Server {
	return &Server{authService: authService, returnAPIData: returnAPIData}
}

func (server *Server) Start() {
	ginEngine := gin.Default()

	ginEngine.GET("/", server.welcome)
	ginEngine.GET("/ping", server.ping)
	ginEngine.POST("/login", server.login)
	ginEngine.POST("/register", server.register)
	ginEngine.POST("/logs", server.getLogs)
	ginEngine.PUT("/users/:id", server.update)
	ginEngine.DELETE("/users/:id", server.delete)
	ginEngine.GET("/refresh_token/:password", server.getRefreshToken)
	ginEngine.GET("/access_token/:password/:refreshtoken", server.getAccessToken)
	ginEngine.GET("/access_token/check/:token", server.checkAccessToken)
	ginEngine.GET("/refresh_token/check/:token", server.checkRefreshToken)
	ginEngine.GET("/product/:barcode", server.retrieveProduct)
	err := ginEngine.Run()
	if err != nil {
		return
	}
}

func (server *Server) retrieveProduct(context *gin.Context) {
	var barcode = context.Param("barcode")
	var productRepo = *server.openFoodFactsService.ProductRepo
	nutrient, err := productRepo.GetProductByBarCode(barcode)
	if nutrient == (entities.Nutrient{}) {
		//CALL OPENFOODFACTS
		//SAVE PRODUCT
	}
	//SEND THE PRODUCT TO FRONTEND
}

func (server *Server) getLogs(context *gin.Context) {
	var logs []entities.UserLogs
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
	var login = entities.Login_info{}
	//binds Json Body to Entities.Login_info Class
	if err := context.ShouldBindJSON(&login); err != nil {
		context.JSON(http.StatusBadRequest, server.returnAPIData.Error(http.StatusBadRequest, err.Error()))
	}
	var userRepo = *server.authService.UserRepo
	loggedIn, err := userRepo.CheckLogin(&login)
	if err != nil {
		context.JSON(http.StatusUnauthorized, server.returnAPIData.Error(http.StatusUnauthorized, "Informations de connexion non valides"))
	} else if loggedIn {
		var user entities.User
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
	var user = entities.User{}
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

func (server *Server) update(context *gin.Context) {
	// Implementation here
}

func (server *Server) delete(context *gin.Context) {
	// Implementation here
}
