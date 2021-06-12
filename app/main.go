package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"i-moscow-backend/app/config"
	"i-moscow-backend/app/db"
	"i-moscow-backend/app/handlers/events"
	"i-moscow-backend/app/handlers/user"
)

func main() {

	config.Load()
	db.Connect()

	app := gin.Default()
	gin.SetMode(gin.DebugMode)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "not found"})
	})

	// user
	app.POST("/auth", user.Auth)
	app.POST("/user", user.Register)
	app.PUT("/user", user.Update)
	app.GET("/user", user.Me)
	app.GET("/user/events", user.UserEvents)
	app.POST("/user/event", user.RegisterToEvent)

	// other
	app.GET("/events", events.GetEvents)
	app.GET("/event/:id", events.GetEvent)

	err := app.Run("localhost:" + config.Port)
	if err != nil {
		fmt.Println("Error in launching backend: " + err.Error())
	}
}
