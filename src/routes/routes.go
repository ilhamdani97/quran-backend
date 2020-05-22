package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	}
}
