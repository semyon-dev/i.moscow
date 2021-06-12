package events

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"i-moscow-backend/app/db"
	"net/http"
)

func GetEvents(c *gin.Context) {

	events := db.GetEvents()

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"events":  events,
	})
}

func GetEvent(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "id can't be empty",
		})
		return
	}

	objectId, _ := primitive.ObjectIDFromHex(id)

	event, _ := db.GetEventByID(objectId)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"event":   event,
	})
}
