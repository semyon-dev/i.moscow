package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"i-moscow-backend/app/db"
	"net/http"
)

func RegisterToEvent(c *gin.Context) {

	jsonInput := struct {
		EventID string `json:"eventId" bson:"eventId"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	email, done := ParseBearer(c)
	if !done {
		return
	}

	user, ok := db.FindUserByEmail(email)
	if ok {
		objID, err := primitive.ObjectIDFromHex(jsonInput.EventID)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		db.AddRegisteredEventToUser(user.Email, objID)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
