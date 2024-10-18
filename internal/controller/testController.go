package controller

import (
	"ginTemplate/internal/constants"
	"ginTemplate/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController interface {
	Test(c *gin.Context)
}

type TestControllerImpl struct {
	testService service.TestService
}

func TestControllerInit(testService service.TestService) *TestControllerImpl {
	return &TestControllerImpl{
		testService: testService,
	}
}

func (u *TestControllerImpl) Test(c *gin.Context) {
	test, err := u.testService.Test(c)
	if err != nil {
		constants.HandleError(c, err)
		return
	}
	c.String(http.StatusOK, test)
	return
}
