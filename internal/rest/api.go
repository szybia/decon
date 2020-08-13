package rest

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateAndRun initialiases and runs the REST API
func CreateAndRun() {
	r := createRouter()

	CreateRoutes(r)

	r.Run()
}

func createRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	gin.DisableConsoleColor()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s ip='%s' method='%s' path='%s' body_size=%d protocol='%s' status_code=%d latency=%d user_agent='%s' error_message='%s'\n",
			param.TimeStamp.Format(time.RFC3339),
			param.ClientIP,
			param.Method,
			param.Path,
			param.BodySize,
			param.Request.Proto,
			param.StatusCode,
			param.Latency.Microseconds(),
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	return r
}
