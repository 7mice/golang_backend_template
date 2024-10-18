package router

import (
	"ginTemplate/app/di"
	"ginTemplate/app/router/routers"
	"ginTemplate/pkg/logging"
	"github.com/gin-gonic/gin"
)

func Init(u *di.Initialization) *gin.Engine {
	router := gin.New()
	router.Use(logging.DefaultStructuredLogger())
	router.Use(gin.Recovery())
	router.Use(u.Middlewares.CorsMiddleware())
	{
		routers.SetupMainGroup(router, u)
	}
	return router
}
