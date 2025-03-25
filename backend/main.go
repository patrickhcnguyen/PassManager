package main

// this is our main server, and is where we use all our routes similar to an index.js on a node server
import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patrickhcnguyen/PassManager/backend/database"
)

func main() {
	// initialize database
	_, err := database.Initialize()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "hello world"})
		})
	}
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })
	router.Run(":8080")
}
