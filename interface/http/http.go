package http

import (
	"context"
	"net/http"

	"template/service/interface/http/api"
	"template/service/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	ROUTE_GET  = "GET"
	ROUTE_POST = "POST"
)

var HttpServer http.Server
var router *gin.Engine
var methodRoutes map[string]map[string]gin.HandlerFunc

func InitHttpServer() {
	// Config Port and Address
	httpPort := viper.GetString("Interface.Http.Port")
	logger.Logger.Infof("HTTP server is initilized")

	// Init Gin Server
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()

	// Config Middleware
	router.Use(SetRequestId())
	router.Use(Logger())
	router.Use(gin.Recovery())

	// Config Path
	api.AddRoute(router)
	AddRoute(router)

	// Init Http Server
	HttpServer = http.Server{
		Addr:    ":" + httpPort,
		Handler: router,
	}

	// Start Server
	logger.Logger.Infof("Serving HTTP API at http://127.0.0.1:%s", httpPort)
	if err := HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Logger.Infof("HTTP server listen and serves failed")
	}

}

func ShutdownHttpServer() {
	logger.Logger.Infof("HTTP server is shutting down")
	if err := HttpServer.Shutdown(context.Background()); err != nil {
		logger.Logger.Infof("HTTP server shut down failed: %s", err)
		return
	}
	logger.Logger.Infof("HTTP server shut down completed")
}

func AddRoute(router *gin.Engine) {
	for method, routes := range methodRoutes {
		if method == ROUTE_GET {
			for routeName, routeFunc := range routes {
				router.GET(routeName, routeFunc)
			}
		} else if method == ROUTE_POST {
			for routeName, routeFunc := range routes {
				router.POST(routeName, routeFunc)
			}
		}
	}
}

func init() {
	methodRoutes = make(map[string]map[string]gin.HandlerFunc)
	methodRoutes[ROUTE_GET] = make(map[string]gin.HandlerFunc)
	methodRoutes[ROUTE_GET]["/ping"] = Ping

	methodRoutes[ROUTE_POST] = make(map[string]gin.HandlerFunc)
}
