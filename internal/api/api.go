package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/api/handlers/v1"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/pkgs/log"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/service"
)

type Options struct {
	Service service.ServiceI
	Log     log.Log
}

func Api(opt Options) *gin.Engine {

	h := v1.NewHandler(opt.Service, opt.Log)

	engine := gin.Default()
	engine.SetTrustedProxies([]string{"192.168.1.1", "10.0.0.0/8"})

	// API Versioning
	apiGroup := engine.Group("/api/v1")
	{
		apiGroup.GET("/ping", h.Ping)

		// Todo routes
		todo := apiGroup.Group("/todo")
		{
			todo.POST("/create", h.CreateTodo)
		}
	}

	return engine
}
