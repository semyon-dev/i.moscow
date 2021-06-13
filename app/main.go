package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"i-moscow-backend/app/config"
	"i-moscow-backend/app/db"
	"i-moscow-backend/app/handlers/events"
	"i-moscow-backend/app/handlers/projects"
	"i-moscow-backend/app/handlers/user"
	"i-moscow-backend/app/handlers/util"
	"i-moscow-backend/app/notifications"
)

func main() {

	config.Load()
	db.Connect()

	notifications.Init()

	app := gin.Default()
	app.Use(cors.Default())

	gin.SetMode(gin.DebugMode)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "not found"})
	})

	// users
	app.POST("/auth", user.Auth)
	app.POST("/user", user.Register)
	app.GET("/auto-completion/:name", util.AutoCompletion)
	app.PUT("/user", user.Update)
	app.GET("/user", user.GetUser)
	app.GET("/user/:id", user.GetUserById)

	// events
	app.GET("/events", events.GetEvents)
	app.GET("/event/:id", events.GetEvent)
	app.GET("/user/events", user.GetUserEvents)
	app.POST("/user/event", user.RegisterToEvent)

	// projects
	app.POST("/project", projects.CreateProject)                                    // new project
	app.GET("/projects", projects.GetProjects)                                      // get projects for users
	app.GET("/projects/my", projects.GetMyProjects)                                 // get my projects
	app.PUT("/project/:id", projects.UpdateProject)                                 // update project info
	app.GET("/project/:id", projects.GetProject)                                    // about project (get by id)
	app.DELETE("/project/:id", projects.DeleteProject)                              // delete project
	app.GET("/project/:id/requests", projects.GetRequests)                          // get requests for projects (only for capitan)
	app.PUT("/project/:id/add-member/:memberId", projects.AddMember)                // add member to project (only for capitan)
	app.DELETE("/project/:id/delete-member/:memberId", projects.DeleteMember)       // delete member from project (only for capitan)
	app.POST("/project/:id/add-request", projects.AddRequest)                       // request participation
	app.DELETE("/project/:id/delete-request/:userId", projects.DeleteRequestMember) // delete request from project (only for capitan)

	err := app.Run("localhost:" + config.Port)
	if err != nil {
		fmt.Println("Error in launching backend: " + err.Error())
	}
}
