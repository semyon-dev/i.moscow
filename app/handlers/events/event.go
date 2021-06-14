package events

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"i-moscow-backend/app/db"
	"net/http"
	"sort"
	"time"
)

func GetEvents(c *gin.Context) {

	events := db.GetEvents()

	sort.Slice(events, func(i, j int) bool {
		return time.Unix(events[i].Date, 0).Unix() < time.Unix(events[j].Date, 0).Unix()
	})

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
