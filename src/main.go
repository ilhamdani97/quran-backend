package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/saefullohmaslul/quran-backend/src/apps"
)

func main() {
	r := gin.Default()

	app := apps.NewApplication(r)
	app.Create()

	if err := godotenv.Load(".env"); err != nil {
		logrus.Error("ENV", err)
	}

	port, found := os.LookupEnv("SERVING_PORT")
	if !found {
		logrus.Warn("Port not found, using default port 8080")
		port = ":8080"
	}

	if err := r.Run(port); err != nil {
		logrus.Error("Error running server", err)
	}
}
