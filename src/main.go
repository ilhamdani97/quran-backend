package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/saefullohmaslul/quran-backend/src/apps"
)

func main() {
	r := gin.Default()

	app := apps.NewApplication(r)
	app.Create()

	if err := r.Run(); err != nil {
		logrus.Error("Error running server", err)
	}
}
