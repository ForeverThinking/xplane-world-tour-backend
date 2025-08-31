package main

import (
	"github.com/ForeverThinking/xplane-world-tour-backend/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.Run(":8080")
}