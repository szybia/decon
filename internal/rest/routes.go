package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var registeredRoutes gin.RoutesInfo

type responseRoute struct {
	Href  string   `json:"href"`
	Verbs []string `json:"verbs"`
}

type responseRoutes map[string]responseRoute

// CreateRoutes adds the routes to the router
func CreateRoutes(router *gin.Engine) {
	//	GET routes
	router.GET("/", getRoot)

	//	Store routes for root  "GET /"  request to list all routes
	registeredRoutes = router.Routes()
}

func getRoot(c *gin.Context) {
	routesResponse := make(responseRoutes, len(registeredRoutes))
	for _, v := range registeredRoutes {
		routesResponse[v.Path] = responseRoute{
			Href:  fmt.Sprintf("%v%v", c.Request.Host, v.Path),
			Verbs: []string{v.Method},
		}
	}

	c.JSON(http.StatusOK, routesResponse)
}
