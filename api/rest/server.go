package rest

import (
	"api/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServerRest struct {
	httpServer *http.Server
	engine     *gin.Engine
	config     *config.Config
}
