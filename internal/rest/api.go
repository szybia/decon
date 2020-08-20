package rest

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/szybia/decon/internal/redis"
	rest "github.com/szybia/decon/internal/rest/routes"
)

var (
	errEndpointNotConfigured = errors.New("\"endpoint\" configuration value not set")
)

// CreateAndRun initialiases and runs the REST API
func CreateAndRun() {
	//	Setup configuration
	//	`viper.Get("redis_addr") => env.get("DECON_REDIS_ADDR")`
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("decon")

	//	Initialise integrations
	err := redis.CreatePool()
	if err != nil {
		die(err)
	}

	//	Create and initialise Router
	r := createRouter()
	rest.CreateRoutes(r)

	//	Run server
	endpoint := viper.GetString("endpoint")
	if endpoint == "" {
		die(errEndpointNotConfigured)
	}

	err = r.Run(endpoint)
	if err != nil {
		die(err)
	}
}

func createRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		//	Don't log meaningless requests
		switch param.Path {
		case "/", "/health":
			return ""
		}
		return fmt.Sprintf("%s ip='%s' method='%s' path='%s' body_size=%d protocol='%s' status_code=%d latency_ns=%d user_agent='%s' error_message='%s'\n",
			param.TimeStamp.Format(time.RFC3339),
			param.ClientIP,
			param.Method,
			param.Path,
			param.BodySize,
			param.Request.Proto,
			param.StatusCode,
			param.Latency.Nanoseconds(),
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	return r
}

func die(e error) {
	fmt.Fprintln(os.Stderr, e)
	os.Exit(1)
}
