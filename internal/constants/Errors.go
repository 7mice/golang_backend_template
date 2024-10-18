package constants

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	InternalError     = errors.New("internal error")
	NotFoundError     = errors.New("resource not found")
	BadRequestError   = errors.New("bad request")
	UnauthorizedError = errors.New("unauthorized")
	ForbiddenError    = errors.New("forbidden")
	ConflictError     = errors.New("conflict")
)

func HandleError(c *gin.Context, err error) {
	switch err {
	case InternalError:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	case NotFoundError:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case BadRequestError:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case UnauthorizedError:
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an unexpected error occurred"})
	}
	c.Abort()
}
