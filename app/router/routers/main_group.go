package routers

import (
	"ginTemplate/app/di"
	"github.com/gin-gonic/gin"
)

func SetupMainGroup(router *gin.Engine, u *di.Initialization) {
	api := router.Group("/api/v1")
	{
		api.GET("/test", u.TestController.Test)
	}
}
