package main

// this is our main server, and is where we use all our routes similar to an index.js on a node server
import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patrickhcnguyen/PassManager/backend/database"
	"github.com/patrickhcnguyen/PassManager/backend/middleware/auth"
	"github.com/patrickhcnguyen/PassManager/backend/routes/userAuth"
)

func main() {
	// initialize database
	_, err := database.Initialize()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	router := gin.Default()

	// cors probably want to move to middleware directory
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cookie")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	api := router.Group("/api")
	{
		// don't require auth
		api.POST("/register", userAuth.Register)
		api.POST("/login", userAuth.Login)

		// require auth
		protected := api.Group("/")
		protected.Use(auth.AuthMiddleware())
		{
			protected.GET("/hello", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{"msg": "hello world"})
			})
			protected.GET("/user", auth.UserHandler)
		}
	}
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })
	router.Run(":8080")
}
