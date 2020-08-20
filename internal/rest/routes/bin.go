package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func putBin(c *gin.Context) {
	key := c.Param("key")
	c.JSON(http.StatusOK, gin.H{
		"key": key,
	})
}
