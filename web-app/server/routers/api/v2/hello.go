package v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stephenwu2020/fabric-cases/web-app/server/pkg/app"
)

// Hello test
func Hello(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, "success", map[string]interface{}{"msg": "Hello, api v2."})
}
