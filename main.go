package main

import (
	"github.com/ForeverThinking/xplane-world-tour-backend/db"
	"github.com/ForeverThinking/xplane-world-tour-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
