package routes

import (
	"mini-task/models"
	"mini-task/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func RegisterUserRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
	v1.GET("/users", )
	v1.POST("/users", )
	v1.PUT("/users", )
	v1.DELETE("/users", )
	}
 }


 func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := storage.CreateUser(user)
	c.JSON(http.StatusCreated, created)
}


func GetUsers(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "5")
	afterStr := c.DefaultQuery("after_id", "0")

	limit, _ := strconv.Atoi(limitStr)
	afterID, _ := strconv.Atoi(afterStr)

	users, nextAfter := storage.GetUsersFiltered(afterID, limit)

	c.JSON(http.StatusOK, gin.H{
		"users":      users,
		"next_after": nextAfter,
	})
}


func UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, ok := storage.UpdateUser(id, user)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	ok := storage.DeleteUser(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}