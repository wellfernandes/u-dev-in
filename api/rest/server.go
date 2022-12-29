package rest

import (
	"api/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServerRest struct {
	HttpServer  *http.Server
	Engine      *gin.Engine
	Config      *config.Config
	Controllers *Controllers
}

type userController interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type Controllers struct {
	UserController userController
}

func NewServerRest(config *config.Config, controllers *Controllers) *ServerRest {

	engine := gin.Default()

	return &ServerRest{
		Engine:      engine,
		Config:      config,
		Controllers: controllers,
	}
}
