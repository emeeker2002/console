package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/open-amt-cloud-toolkit/console/pkg/consoleerrors"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(c *gin.Context, err error) {
	switch err.(type) {
	case validator.ValidationErrors:
		c.AbortWithStatusJSON(http.StatusBadRequest, response{err.Error()})
	case consoleerrors.NotFoundError:
		c.AbortWithStatusJSON(http.StatusNotFound, response{err.Error()})
	case consoleerrors.NotUniqueError:
		c.AbortWithStatusJSON(http.StatusBadRequest, response{err.Error()})
	case consoleerrors.DatabaseError:
		c.AbortWithStatusJSON(http.StatusInternalServerError, response{err.Error()}) // Update error message to err.Message with either AMT or Database error
	case consoleerrors.AMTError:
		c.AbortWithStatusJSON(http.StatusInternalServerError, response{err.Error()})

	// case consoleerrors.ConsoleError:
	// 	switch origErr := errTyped.OriginalError.(type) {
	// 	case consoleerrors.NotFoundError:
	// 		c.AbortWithStatusJSON(http.StatusNotFound, response{origErr.Error()})
	// 	case consoleerrors.NotUniqueError:
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, response{origErr.Error()})
	// 	case consoleerrors.DatabaseError:
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, response{origErr.Error()}) // Update error message to origErr.Message with either AMT or Database error
	// 	case consoleerrors.AMTError:
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, response{origErr.Error()})
	// 	default:
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, response{"general error"})
	// 	}
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, response{"general error"})
	}
}
