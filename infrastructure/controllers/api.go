package api

import (
	"gaia-api/domain/entities"
	"gaia-api/domain/services"
	"net/http"

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

	gin_engine.Run()
}

func (server *Server) welcome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Title": "Gaia"})
}
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
		context.JSON(http.StatusInternalServerError, err.Error())
	} else if logged_in {
		//a function that generates a token using JWT
		context.JSON(http.StatusAccepted, gin.H{"cnx_Token": "token"})
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
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
