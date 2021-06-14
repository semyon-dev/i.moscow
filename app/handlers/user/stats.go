package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"i-moscow-backend/app/db"
	"net/http"
)

func IncreaseStat(c *gin.Context) {
	statName := c.Param("statName")
	res, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	err = db.IncreaseStat(res, statName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
