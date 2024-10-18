package repository

import (
	"ginTemplate/internal/constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TestRepository interface {
	Test(c *gin.Context) (string, error)
}

type TestRepositoryImpl struct {
	db *gorm.DB
}

func TestRepositoryInit(db *gorm.DB) *TestRepositoryImpl {
	return &TestRepositoryImpl{
		db: db,
	}
}

func (u *TestRepositoryImpl) Test(c *gin.Context) (string, error) {
	return "test", constants.UnauthorizedError
}
