package service

import (
	"ginTemplate/internal/repository"
	"github.com/gin-gonic/gin"
)

type TestService interface {
	Test(c *gin.Context) string
}

type TestServiceImpl struct {
	testRepository repository.TestRepository
}

func TestServiceInit(testRepository repository.TestRepository) *TestServiceImpl {
	return &TestServiceImpl{
		testRepository: testRepository,
	}
}

func (u *TestServiceImpl) Test(c *gin.Context) string {
	return u.testRepository.Test(c)
}
