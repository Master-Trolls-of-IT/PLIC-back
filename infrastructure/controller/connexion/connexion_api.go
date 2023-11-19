package connexion

import (
	interfaces "gaia-api/application/interface"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Connexion struct {
	ginEngine     *gin.Engine
	returnAPIData *interfaces.ReturnAPIData
}

func NewConnexionController(ginEngine *gin.Engine, returnAPIData *interfaces.ReturnAPIData) *Connexion {
	connexion := &Connexion{ginEngine: ginEngine, returnAPIData: returnAPIData}
	connexion.Start()
	return connexion
}

func (connexion *Connexion) Start() {
	connexion.ginEngine.GET("/", connexion.welcome)
	connexion.ginEngine.GET("/ping", connexion.ping)
}

func (connexion *Connexion) welcome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Title": "Gaia"})
}

func (connexion *Connexion) ping(context *gin.Context) {
	context.JSON(http.StatusOK, connexion.returnAPIData.Ping())
}
