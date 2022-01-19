package controller

import (
	"my_kanban_board/database"
	"my_kanban_board/helper"
	"my_kanban_board/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TaskPost(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	Task := model.Task{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		c.ShouldBind(&Task)
	}

	Task.Status = false
	Task.UserID = userID

	err := db.Debug().Create(&Task).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          Task.ID,
		"title":       Task.Title,
		"status":      Task.Status,
		"description": Task.Description,
		"user_id":     Task.UserID,
		"category_id": Task.CategoryID,
		"Created_at":  Task.CreatedAt,
	})
}

func TaskView(c *gin.Context) {
	db := database.GetDB()
	Task := []model.Task{}

	err := db.Find(&Task)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error,
		})
		return
	}

	getTask := []model.GetTask{}
	for i := 0; i < len(Task); i++ {
		tempGetTask := model.GetTask{}
		User := model.User{}
		err := db.Where("id = ?", Task[i].UserID).First(&User)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error,
			})
			return
		}
		tempGetTask.ID = Task[i].ID
		tempGetTask.Title = Task[i].Title
		tempGetTask.Status = Task[i].Status
		tempGetTask.Description = Task[i].Description
		tempGetTask.UserID = Task[i].UserID
		tempGetTask.CategoryID = Task[i].CategoryID
		tempGetTask.CreatedAt = Task[i].CreatedAt
		tempGetTask.User.ID = User.ID
		tempGetTask.User.Email = User.Email
		tempGetTask.User.FullName = User.FullName

		getTask = append(getTask, tempGetTask)
	}

	c.JSON(http.StatusOK, getTask)

}

func TaskUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)

	Task := model.Task{}
	taskID, _ := strconv.Atoi(c.Param("taskId"))
	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		c.ShouldBind(&Task)
	}

	NewTask := model.Task{}
	_ = db.Where("id = ?", taskID).First(&NewTask)

	NewTask.Title = Task.Title
	NewTask.Description = Task.Description

	err := db.Model(&NewTask).Where("id=?", taskID).Updates(
		model.Task{
			Title:       NewTask.Title,
			Description: NewTask.Description,
		}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          NewTask.ID,
		"title":       NewTask.Title,
		"description": NewTask.Description,
		"status":      NewTask.Status,
		"user_id":     NewTask.UserID,
		"category_id": NewTask.CategoryID,
		"updated_at":  NewTask.UpdatedAt,
	})

}

func TaskUpdateStatus(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)

	Task := model.Task{}
	taskID, _ := strconv.Atoi(c.Param("taskId"))
	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		c.ShouldBind(&Task)
	}

	NewTask := model.Task{}
	_ = db.Where("id = ?", taskID).First(&NewTask)
	NewTask.Status = Task.Status

	db.Model(&model.Task{}).Where("id=?", taskID).Update("status", NewTask.Status)

	c.JSON(http.StatusOK, gin.H{
		"id":          NewTask.ID,
		"title":       NewTask.Title,
		"description": NewTask.Description,
		"status":      NewTask.Status,
		"user_id":     NewTask.UserID,
		"category_id": NewTask.CategoryID,
		"updated_at":  NewTask.UpdatedAt,
	})
}

func TaskUpdateCategory(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)

	Task := model.Task{}
	taskID, _ := strconv.Atoi(c.Param("taskId"))
	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		c.ShouldBind(&Task)
	}

	NewTask := model.Task{}
	_ = db.Where("id = ?", taskID).First(&NewTask)
	NewTask.Status = Task.Status

	db.Model(&model.Task{}).Where("id=?", taskID).Update("category_id", NewTask.CategoryID)

	c.JSON(http.StatusOK, gin.H{
		"id":          NewTask.ID,
		"title":       NewTask.Title,
		"description": NewTask.Description,
		"status":      NewTask.Status,
		"user_id":     NewTask.UserID,
		"category_id": NewTask.CategoryID,
		"updated_at":  NewTask.UpdatedAt,
	})
}

func TaskDelete(c *gin.Context) {
	db := database.GetDB()
	Task := model.Task{}
	taskID, _ := strconv.Atoi(c.Param("taskId"))
	db.Delete(&Task, taskID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Task has been successfully deleted",
	})

}
