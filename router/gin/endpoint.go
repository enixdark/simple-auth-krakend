package gin

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/proxy"
	krakendgin "github.com/devopsfaith/krakend/router/gin"
	"strings"
)

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
}

// HandlerFactory decorates a krakendgin.HandlerFactory with the auth layer
func HandlerFactory(hf krakendgin.HandlerFactory) krakendgin.HandlerFactory {
	return func(configuration *config.EndpointConfig, proxy proxy.Proxy) gin.HandlerFunc {
		next := hf(configuration, proxy)

		return func(c *gin.Context) {

			auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

			if len(auth) != 2 || auth[0] != "Basic" {
				respondWithError(401, "Unauthorized", c)
				return
			}
			payload, _ := base64.StdEncoding.DecodeString(auth[1])
			pair := strings.SplitN(string(payload), ":", 2)
			if len(pair) != 2 {
				respondWithError(401, "Unauthorized", c)
				return
			}
			next(c)
		}
	}
}
