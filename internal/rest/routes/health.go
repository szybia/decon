package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/szybia/decon/internal/redis"
)

func health(c *gin.Context) {
	conn := redis.GetConnection()
	defer conn.Close()

	//	Check Redis connection
	_, err := conn.Do("PING")
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"ready": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ready": true,
	})
}
