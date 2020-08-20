package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseRoute struct {
	Href  string   `json:"href"`
	Verbs []string `json:"verbs"`
}

type responseRoutes map[string]responseRoute

var registeredRoutes gin.RoutesInfo

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
