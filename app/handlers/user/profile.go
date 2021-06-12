package user

import (
	"github.com/gin-gonic/gin"
	"i-moscow-backend/app/db"
	"net/http"
)

func Update(c *gin.Context) {
	// TODO
}

func Me(c *gin.Context) {
	username, done := ParseBearer(c)
	if done {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	user, ok := db.FindUserByEmail(username)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
		return
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}

func UserEvents(c *gin.Context) {
	username, done := ParseBearer(c)
	if done {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	user, ok := db.FindUserByEmail(username)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"userEvents": user.RegisteredEvents,
		})
		return
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}
