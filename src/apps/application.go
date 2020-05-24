package apps

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/saefullohmaslul/quran-backend/src/routes"
)

// Application is initialize app with gin
type Application struct {
	Route *gin.Engine
}

// NewApplication is constructor for app
func NewApplication(route *gin.Engine) *Application {
	return &Application{Route: route}
}

// Create is method to create application
func (a Application) Create() {
	configureEndpoint(a.Route)
}

func configureEndpoint(r *gin.Engine) {
	g := r.Group("/api/v1")

	routes.Router(g)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Route not found",
			"data":    nil,
			"errors": []map[string]interface{}{
				gin.H{
					"flag":    "ROUTE_NOT_FOUND",
					"message": "The route you are looking for is not found",
				},
			},
		})
	})
}
