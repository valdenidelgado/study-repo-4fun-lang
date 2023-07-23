package main

import (
	"api/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.AppRoutes(app)

	app.Run(":8080")
}
