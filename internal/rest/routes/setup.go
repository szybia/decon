package rest

import (
	"github.com/gin-gonic/gin"
)

// CreateRoutes adds the routes to the router
func CreateRoutes(router *gin.Engine) {
	//	404
	router.NoRoute(noRoute)

	//	GET routes
	router.GET("/", getRoot)
	router.GET("/health", health)

	//	PUT routes
	router.PUT("/bin/:key", putBin)

	//	Store routes for root  "GET /"  request to list all routes
	registeredRoutes = router.Routes()
}
