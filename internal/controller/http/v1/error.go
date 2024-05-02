package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
	"github.com/open-amt-cloud-toolkit/console/pkg/consoleerrors"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(c *gin.Context, err error) {
	if errors.Is(err, consoleerrors.ErrNotFound) {
		msg := err.Error()
		c.AbortWithStatusJSON(http.StatusNotFound, response{msg})
		return
	} else if _, ok := err.(validator.ValidationErrors); ok {
		msg := err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, response{msg})
	} else if errors.Is(err, consoleerrors.ErrNotUnique) {
		msg := err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, response{msg})
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response{err.Error()})
	}
}
