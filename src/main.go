package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()

	if err := r.Run(); err != nil {
		logrus.Error("Error running server", err)
	}
}
