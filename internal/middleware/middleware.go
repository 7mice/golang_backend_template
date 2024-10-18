package middleware

import (
	"ginTemplate/internal/constants"
	"ginTemplate/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Middlewares interface {
	CorsMiddleware() gin.HandlerFunc
	TestMiddleware() gin.HandlerFunc
}

type MiddlewaresImpl struct {
	testService service.TestService
}

func MiddlewaresInit() *MiddlewaresImpl {
	return &MiddlewaresImpl{}
}

func (u *MiddlewaresImpl) TestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		{
			constants.HandleError(c, constants.BadRequestError)
		}
		c.Next()
	}
}

func (u *MiddlewaresImpl) CorsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	return cors.New(config)
}
