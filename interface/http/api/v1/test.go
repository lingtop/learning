package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetName(context *gin.Context) {
	contextLogger, _ := context.Get("logger")
	apiLogger := contextLogger.(*zap.SugaredLogger)
	apiLogger.Info("call interface get name")

	context.JSON(http.StatusOK, gin.H{
		"test": "1",
	})
}

func init() {
	methodRoutes[ROUTE_GET]["/get-name"] = GetName
}
