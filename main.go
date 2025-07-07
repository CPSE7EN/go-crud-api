package main

import (
	"github.com/CPSE7EN/go-crud-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}
