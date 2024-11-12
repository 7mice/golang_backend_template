package repository

import (
	"ginTemplate/internal/constants"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type TestRepository interface {
	Test(c *gin.Context) (string, error)
}

type TestRepositoryImpl struct {
	db *sqlx.DB
}

func TestRepositoryInit(db *sqlx.DB) *TestRepositoryImpl {
	return &TestRepositoryImpl{
		db: db,
	}
}

func (u *TestRepositoryImpl) Test(c *gin.Context) (string, error) {
	return "test", constants.UnauthorizedError
}
