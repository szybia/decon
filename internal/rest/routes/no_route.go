package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func noRoute(c *gin.Context) {
	resp := errorResponse{
		Errors: []responseError{
			pageNotFoundError,
		},
	}

	c.JSON(http.StatusNotFound, resp)
}
