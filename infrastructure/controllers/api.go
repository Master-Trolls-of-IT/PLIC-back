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

// Function that checks if the access token is valid, it takes an access token as parameter and returns a JSON with this structure: {"valid": true}
func (server *Server) checkAccessToken(context *gin.Context) {
	retVal, err := verifyAccessToken(context.Param("token"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"valid": retVal})
	}
}

// Function that checks if the refresh token is valid, it takes a refresh token as parameter and returns a JSON with this structure: {"valid": true}
func (server *Server) checkRefreshToken(context *gin.Context) {
	retVal, err := verifyRefreshToken(context.Param("token"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"valid": retVal})
	}
}

// Function that generates an access token, it takes a password and a refreshtoken as parameter and returns a JSON with this structure : { "token": generatedToken }
func (server *Server) getAccessToken(context *gin.Context) {
	// Use GenerateAccessToken function to generate a new access token
	accessToken, err := generateAccessToken(context.Param("password"), []byte(context.Param("refreshtoken")))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"token": accessToken})
}

// Function that generates a refresh token, it takes a password as parameter and returns a JSON with this structure : { "token": generatedToken }
func (server *Server) getRefreshToken(context *gin.Context) {
	// Use GenerateRefreshToken function to generate a new refresh token
	refreshToken, err := generateRefreshToken([]byte(context.Param("password")))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"token": refreshToken})
}

// Welcome function that returns a JSON with this structure : { "Title": "Gaia" }
func (server *Server) welcome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Title": "Gaia"})
}

// Ping function that returns a JSON with this structure : { "ping": "pong" }
func (server *Server) ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

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
