package projects

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"i-moscow-backend/app/db"
	"i-moscow-backend/app/model"
	"i-moscow-backend/app/notifications"
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
	project.TeamIDs = make([]primitive.ObjectID, 0, 0)
	project.Skills = make([]string, 0, 0)
	project.RequestedIds = make([]primitive.ObjectID, 0, 0)
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	project.TeamCapitan, _ = primitive.ObjectIDFromHex(id)
	project.Id = primitive.NewObjectID()
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
	id, _, _ := session.ParseBearer(c)
	idObject, _ := primitive.ObjectIDFromHex(id)
	projects := db.GetProjects(idObject)
	c.JSON(http.StatusOK, gin.H{"message": "ok", "projects": projects})
}

func GetProject(c *gin.Context) {
	id := c.Param("id")
	idO, _ := primitive.ObjectIDFromHex(id)
	project, _ := db.GetProjectById(idO)
	c.JSON(http.StatusOK, gin.H{"message": "ok", "project": project})
}

func GetRequests(c *gin.Context) {
	id := c.Param("id")
	idO, _ := primitive.ObjectIDFromHex(id)
	project, _ := db.GetProjectById(idO)
	c.JSON(http.StatusOK, gin.H{"message": "ok", "requestedUsers": project.RequestedIds})
}

func AddRequest(c *gin.Context) {
	userId, _, done := session.ParseBearer(c)
	if !done {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid token",
		})
		return
	}
	projectId := c.Param("id")
	userIdObjectId, _ := primitive.ObjectIDFromHex(userId)
	projectIdObectId, _ := primitive.ObjectIDFromHex(projectId)
	err := db.AddRequestMemberToProject(projectIdObectId, userIdObjectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error or invalid project",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func GetMyProjects(c *gin.Context) {
	id, _, done := session.ParseBearer(c)
	if !done {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid token",
		})
		return
	}
	projects := db.GetMyProjects(id)
	c.JSON(http.StatusOK, gin.H{"message": "ok", "projects": projects})
}

func AddMember(c *gin.Context) {
	projectId := c.Param("id")
	memberId := c.Param("memberId")
	memberIdObjectId, err1 := primitive.ObjectIDFromHex(memberId)
	projectIdObjectId, err2 := primitive.ObjectIDFromHex(projectId)
	if err1 != nil || err2 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad memberId or projectId",
		})
		return
	}
	err := db.AddMemberToProject(projectIdObjectId, memberIdObjectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error or invalid project",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
	user, _ := db.FindUserById(memberIdObjectId)
	project, _ := db.GetProjectById(projectIdObjectId)
	notifications.Send(user.DeviceToken, "Ура! Теперь вы часть проекта.", "Ваша заявка была одобрена в проект \n"+project.Name)
}

func DeleteMember(c *gin.Context) {
	projectId := c.Param("id")
	memberId := c.Param("memberId")
	memberIdObjectId, err1 := primitive.ObjectIDFromHex(memberId)
	projectIdObjectId, err2 := primitive.ObjectIDFromHex(projectId)
	if err1 != nil || err2 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad memberId or projectId",
		})
		return
	}
	err := db.DeleteMemberFromProject(projectIdObjectId, memberIdObjectId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error or invalid project",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func DeleteRequestMember(c *gin.Context) {
	projectId := c.Param("id")
	userId := c.Param("userId")
	userIdIdObjectId, err1 := primitive.ObjectIDFromHex(userId)
	projectIdObjectId, err2 := primitive.ObjectIDFromHex(projectId)
	if err1 != nil || err2 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad userId or projectId",
		})
		return
	}
	err := db.DeleteRequestFromProject(projectIdObjectId, userIdIdObjectId)
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
