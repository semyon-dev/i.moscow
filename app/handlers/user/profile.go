package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"i-moscow-backend/app/db"
	"i-moscow-backend/app/model"
	"i-moscow-backend/app/session"
	"net/http"
)

func Update(c *gin.Context) {
	_, _, done := session.ParseBearer(c)
	if !done {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid token",
		})
		return
	}
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}
	if user.Id.IsZero() {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "user id can't be empty",
		})
		return
	}
	err := db.UpdateUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func GetUser(c *gin.Context) {
	_, email, done := session.ParseBearer(c)
	if !done {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid token",
		})
		return
	}
	user, ok := db.FindUserByEmail(email)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
		return
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	user, ok := db.FindUserById(objectId)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
		return
	}
	c.AbortWithStatus(http.StatusBadRequest)
}

func GetUserEvents(c *gin.Context) {
	id, _, done := session.ParseBearer(c)
	if !done {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	objectId, _ := primitive.ObjectIDFromHex(id)
	user, ok := db.FindUserById(objectId)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"userEvents": user.RegisteredEvents,
		})
		return
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}
