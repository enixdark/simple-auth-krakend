package auth

import (
	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/gin-gonic/gin"
)

// NewEngine creates a new gin engine with some default values and a secure middleware
func NewEngine(cfg config.ServiceConfig, logger logging.Logger) *gin.Engine {
	engine := gin.Default()

	engine.RedirectTrailingSlash = true
	engine.HandleMethodNotAllowed = true

	return engine
}
