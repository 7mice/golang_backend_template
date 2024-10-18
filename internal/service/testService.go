package service

import (
	"ginTemplate/internal/repository"
	"github.com/gin-gonic/gin"
)

type TestService interface {
	Test(c *gin.Context) (string, error)
}

type TestServiceImpl struct {
	testRepository repository.TestRepository
}

func TestServiceInit(testRepository repository.TestRepository) *TestServiceImpl {
	return &TestServiceImpl{
		testRepository: testRepository,
	}
}

func (u *TestServiceImpl) Test(c *gin.Context) (string, error) {
	return u.testRepository.Test(c)
}
