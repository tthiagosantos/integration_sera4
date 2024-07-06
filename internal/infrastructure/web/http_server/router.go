package http_server

import (
	"github.com/gin-gonic/gin"
	"integrations_apis/internal/infrastructure/web/http_server/handlers"
)

type Router struct {
	HealthCheck  *handlers.HealthCheckHandler
	Sera4Handler *handlers.Sera4
}

func NewRouter(healCheckHandler *handlers.HealthCheckHandler, Sera4Handler *handlers.Sera4) *Router {
	return &Router{
		HealthCheck:  healCheckHandler,
		Sera4Handler: Sera4Handler,
	}
}

func (r *Router) Register() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/check", r.HealthCheck.HealthCheck)
		//Sera4
		api.POST("/sera4/user", r.Sera4Handler.CreateUser)
		api.DELETE("/sera4/user/:id", r.Sera4Handler.DeleteUser)
		api.GET("/sera4/user/:id", r.Sera4Handler.GetUser)
		api.POST("sera4/key", r.Sera4Handler.CreateKey)
		api.DELETE("sera4/key/:id", r.Sera4Handler.DeleteKey)
	}

	return router
}
