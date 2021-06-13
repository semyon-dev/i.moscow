package projects

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"i-moscow-backend/app/db"
	"i-moscow-backend/app/model"
	"i-moscow-backend/app/session"
	"net/http"
)

func CreateProject(c *gin.Context) {

	id, _, done := session.ParseBearer(c)
	if !done {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "invalid token",
		})
		return
	}

	var project model.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	project.TeamCapitanID, _ = primitive.ObjectIDFromHex(id)
	err := db.Insert("projects", project)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func DeleteProject(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	err := db.Delete("projects", id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func GetProjects(c *gin.Context) {
	projects := db.Get("projects")
	c.JSON(http.StatusOK, gin.H{"message": "ok", "projects": projects})
}

func GetProject(c *gin.Context) {
	id := c.Param("id")
	project, _ := db.GetProjectById(id)
	c.JSON(http.StatusOK, gin.H{"message": "ok", "project": project})
}

func GetRequests(c *gin.Context) {
	id := c.Param("id")
	project, _ := db.GetProjectById(id)
	c.JSON(http.StatusOK, gin.H{"message": "ok", "requestedUsers": project.RequestedIds})
}

func GetMyProjects(c *gin.Context) {
	id, _, done := session.ParseBearer(c)
	if !done {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "invalid token",
		})
		return
	}
	projects := db.GetProjects(id)
	c.JSON(http.StatusOK, gin.H{"message": "ok", "projects": projects})
}

func AddMember(c *gin.Context) {
	projectId := c.Param("id")
	id, _, done := session.ParseBearer(c)
	if !done {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "invalid token",
		})
		return
	}
	idMember, _ := primitive.ObjectIDFromHex(id)
	err := db.AddMemberToProject(projectId, idMember)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error or invalid project",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func UpdateProject(c *gin.Context) {
	projectId := c.Param("id")
	_, _, done := session.ParseBearer(c)
	if !done {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "invalid token",
		})
		return
	}
	var project model.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}
	project.Id, _ = primitive.ObjectIDFromHex(projectId)
	err := db.UpdateProject(project)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error or invalid project",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
