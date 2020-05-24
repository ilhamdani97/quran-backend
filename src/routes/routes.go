package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router is method to create routing point
func Router(g *gin.RouterGroup) {
	{
		g.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "Testing OK",
				"data":    nil,
				"errors":  nil,
			})
		})

		g.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "Health OK",
				"data": gin.H{
					"database": "OK",
				},
				"errors": nil,
			})
		})
	}
}
