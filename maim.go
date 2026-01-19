package main

import (
	"mini-task/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RegisterUserRoutes(r)
	r.Run(":8080")
}