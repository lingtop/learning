package api

import (
	v1 "template/service/interface/http/api/v1"

	"github.com/gin-gonic/gin"
)

func AddRoute(engine *gin.Engine) {
	apiRoute := engine.Group("/api")
	v1.AddRoute(apiRoute)
}
