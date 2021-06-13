package util

import (
	"github.com/gin-gonic/gin"
	"i-moscow-backend/app/db"
	"net/http"
)

func AutoCompletion(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "name can't be empty",
		})
		return
	}
	reply := db.FullTextSearch(name, 50)
	c.JSON(http.StatusOK, gin.H{
		"message":         "ok",
		"autoCompletions": reply,
	})
}
