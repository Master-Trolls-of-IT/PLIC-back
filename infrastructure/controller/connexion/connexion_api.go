package connexion

import (
	"gaia-api/application/returnAPI"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Connexion struct {
	ginEngine *gin.Engine
}

func NewConnexionController(ginEngine *gin.Engine) *Connexion {
	connexion := &Connexion{ginEngine: ginEngine}
	connexion.Start()
	return connexion
}

func (connexion *Connexion) Start() {
	connexion.ginEngine.GET("/", connexion.welcome)
	connexion.ginEngine.GET("/ping", connexion.ping)
}

func (connexion *Connexion) welcome(context *gin.Context) {
	returnAPI.Success(context, http.StatusOK, gin.H{"Title": "Gaïa: Nature's Kitchen"})
}

func (connexion *Connexion) ping(context *gin.Context) {
	returnAPI.Success(context, http.StatusOK, gin.H{"response": "Pong"})
}
