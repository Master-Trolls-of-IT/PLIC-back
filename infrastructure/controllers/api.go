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

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
package controllers

import (
	"gaia-api/domain/entities"
	"gaia-api/domain/services"
	"net/http"

	_ "github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

type Server struct {
	auth_service *services.Auth_service
}

func NewServer(auth_service *services.Auth_service) *Server {
	return &Server{auth_service: auth_service}
}

func (server *Server) Start() {
	gin_engine := gin.Default()

	gin_engine.GET("/", server.welcome)
	gin_engine.GET("/ping", server.ping)
	gin_engine.POST("/login", server.login)
	gin_engine.POST("/register", server.register)
	gin_engine.PUT("/users/:id", server.update)
	gin_engine.DELETE("/users/:id", server.delete)
	gin_engine.GET("/refresh_token/:password", server.getRefreshToken)
	gin_engine.GET("/access_token/:password/:refreshtoken", server.getAccessToken)
	gin_engine.GET("/access_token/check/:token", server.checkAccessToken)
	gin_engine.GET("/refresh_token/check/:token", server.checkRefreshToken)
	gin_engine.Run()
}

// checkAccessToken godoc
//	@Summary		Checks token validity
//	@Description	Check if access token is valid
//	@Accept			json
//	@Produce		json
//	@Param			token	path		int	true	"Access token"
//	@Success		200	{object}	bool
//	@Router			/access_token/check/{token} [get]
func (server *Server) checkAccessToken(context *gin.Context) {
	retVal, err := verifyAccessToken(context.Param("token"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"valid": retVal})
	}
}

// checkRefreshToken godoc
//	@Summary		Checks token validity
//	@Description	Check if refresh token is valid
//	@Accept			json
//	@Produce		json
//	@Param			token	path		int	true	"Refresh token"
//	@Success		200	{object}	bool
//	@Router			/refresh_token/check/{token} [get]
func (server *Server) checkRefreshToken(context *gin.Context) {
	retVal, err := verifyRefreshToken(context.Param("token"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"valid": retVal})
	}
}

// getAccessToken godoc
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
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"token": accessToken})
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
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"token": refreshToken})
}


type Welcome struct {
	Title string `json:"Title" example:"Gaia"`
}
// welcome godoc
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
// 	@Summary		Ping message
// 	@Description	Checks if server is up
// 	@Accept			json
// 	@Produce		json
// 	@Success		200	{object}	Ping
// 	@Router			/ping [get]
func (server *Server) ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

// login godoc
//	@Summary		Login
//	@Description	Login function that returns a JSON with this structure : { "cnx_Token": "token" }
//	@Accept			json
//	@Produce		json
//	@Param			login	body		object	true	"Login info"
//	@Success		202	{object}	string
//	@Router			/login [post]
func (server *Server) login(context *gin.Context) {
	var login = entities.Login_info{}
	//binds Json Body to Entities.Login_info Class
	if err := context.ShouldBindJSON(&login); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var user_repo = *server.auth_service.User_repo
	logged_in, err := user_repo.CheckLogin(&login)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
	} else if logged_in {
		//a function that generates a token using JWT
		context.JSON(http.StatusAccepted, gin.H{"cnx_Token": "token"})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// register godoc
//	@Summary		Register
//	@Description	Register function that returns a JSON with this structure : { "message": "User registered successfully" }
//	@Accept			json
//	@Produce		json
//	@Param			user	body		object	true	"User info"
//	@Success		200	{object}	string
//	@Router			/register [post]
func (server *Server) register(context *gin.Context) {
	var user = entities.User{}
	//binds Json Body to Entities.User Class
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user_repo = *server.auth_service.User_repo
	registered, err := user_repo.Register(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if registered {
		context.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	} else {
		context.JSON(http.StatusConflict, gin.H{"error": "Username or Email already taken"})
	}
}

func (server *Server) update(context *gin.Context) {
	// Implementation here
}

func (server *Server) delete(context *gin.Context) {
	// Implementation here
}
