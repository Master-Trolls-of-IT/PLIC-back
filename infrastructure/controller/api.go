//	@title			PLIC BACKEND API
//	@version		1.0
//	@description	This is a simple API for PLIC BACKEND
//	@termsOfService	à compléter

//	@contact.name	à compléter
//	@contact.email	gaiank2024@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		plic-back-qp6wugltyq-ew.a.run.app/

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/

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

// checkAccessToken godoc
//
//	@Summary		Checks token validity
//	@Description	Check if access token is valid
//	@Accept			json
//	@Produce		json
//	@Param			token	path		int	true	"Access token"
//	@Success		200	{object}	bool
//	@Router			/access_token/check/{token} [get]
func (server *Server) checkAccessToken(context *gin.Context) {
	isTokenValid, err := verifyAccessToken(context.Param("token"))
	if err != nil {
		context.JSON(http.StatusBadRequest, server.returnAPIData.Error(http.StatusBadRequest, err.Error()))
	} else {
		context.JSON(http.StatusOK, server.returnAPIData.CheckToken(isTokenValid))
	}
}

// checkRefreshToken godoc
//
//	@Summary		Checks token validity
//	@Description	Check if refresh token is valid
//	@Accept			json
//	@Produce		json
//	@Param			token	path		int	true	"Refresh token"
//	@Success		200	{object}	bool
//	@Router			/refresh_token/check/{token} [get]
func (server *Server) checkRefreshToken(context *gin.Context) {
	isTokenValid, err := verifyRefreshToken(context.Param("token"))
	if err != nil {
		context.JSON(http.StatusBadRequest, server.returnAPIData.Error(http.StatusBadRequest, err.Error()))
	} else {
		context.JSON(http.StatusOK, server.returnAPIData.CheckToken(isTokenValid))
	}
}

// getAccessToken godoc
//
//	@Summary		Generates a new access token
//	@Description	Generates a new access token
//	@Accept			json
//	@Produce		json
//	@Param			password	path		string	true	"Hashed User password"
//	@Param			refreshtoken	path		string	true	"Refresh token"
//	@Success		200	{object}	string
//	@Router			/access_token/{password}/{refreshtoken} [get]
func (server *Server) getAccessToken(context *gin.Context) {
	// Use GenerateAccessToken function to generate a new access token
	accessToken, err := generateAccessToken(context.Param("password"), []byte(context.Param("refreshtoken")))
	if err != nil {
		context.JSON(http.StatusBadRequest, server.returnAPIData.Error(http.StatusBadRequest, err.Error()))
	}
	context.JSON(http.StatusOK, server.returnAPIData.GetToken(accessToken))
}

// getRefreshToken godoc
//	@Summary		Generates a new refresh token
//	@Description	Generates a new refresh token
//	@Accept			json
//	@Produce		json
//	@Param			password	path		string	true	"Hashed User password"
//	@Success		200	{object}	string
//	@Router			/refresh_token/{password} [get]

func (server *Server) getRefreshToken(context *gin.Context) {
	// Use GenerateRefreshToken function to generate a new refresh token
	refreshToken, err := generateRefreshToken([]byte(context.Param("password")))
	if err != nil {
		context.JSON(http.StatusBadRequest, server.returnAPIData.Error(http.StatusBadRequest, err.Error()))
	}
	context.JSON(http.StatusOK, server.returnAPIData.GetToken(refreshToken))
}

type Welcome struct {
	Title string `json:"Title" example:"Gaia"`
}

// welcome godoc
//
//	@Summary		Welcome message
//	@Description	Welcome function that returns a JSON with this structure : { "Title": "Gaia" }
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Welcome
//	@Router			/ [get]
func (server *Server) welcome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Title": "Gaia"})
}

type Ping struct {
	Title string `json:"ping" example:"pong"`
}

// ping godoc
//
//	@Summary		Ping message
//	@Description	Checks if server is up
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Ping
//	@Router			/ping [get]
func (server *Server) ping(context *gin.Context) {
	context.JSON(http.StatusOK, server.returnAPIData.Ping())
}

// login godoc
//
//	@Summary		Login
//	@Description	Login function that returns a JSON with this structure : { "cnx_Token": "token" }
//	@Accept			json
//	@Produce		json
//	@Param			login	body		object	true	"Login info"
//	@Success		202	{object}	string
//	@Router			/login [post]
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

// register godoc
//
//	@Summary		Register
//	@Description	Register function that returns a JSON with this structure : { "message": "User registered successfully" }
//	@Accept			json
//	@Produce		json
//	@Param			user	body		object	true	"User info"
//	@Success		200	{object}	string
//	@Router			/register [post]
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
