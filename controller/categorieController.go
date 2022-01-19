package controller

import (
	"my_kanban_board/database"
	"my_kanban_board/helper"
	"my_kanban_board/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CategoriePost(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	Category := model.Category{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Category)
	} else {
		c.ShouldBind(&Category)
	}

	err := db.Debug().Create(&Category).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Category.ID,
		"type":       Category.Type,
		"created_at": Category.CreatedAt,
	})

}

func CategorieViewAll(c *gin.Context) {
	db := database.GetDB()
	Category := []model.Category{}

	err := db.Find(&Category)

	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error,
		})
		return
	}
	getCategory := []model.GetCategory{}

	for i := 0; i < len(Category); i++ {
		categoryID := Category[i].ID
		tempGetCategory := model.GetCategory{}
		Tasks := []model.Task{}

		err := db.Where("category_id=?", categoryID).Find(&Tasks)
		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error,
			})
			return
		}
		tempGetCategory.ID = Category[i].ID
		tempGetCategory.Type = Category[i].Type
		tempGetCategory.UpdatedAt = Category[i].UpdatedAt
		tempGetCategory.CreatedAt = Category[i].CreatedAt
		if len(Tasks) > 0 {
			for j := 0; j < len(Tasks); j++ {
				tempTask := model.CategoryTasks{}
				tempTask.ID = Tasks[j].ID
				tempTask.Title = Tasks[j].Title
				tempTask.Description = Tasks[j].Description
				tempTask.UserID = Tasks[j].UserID
				tempTask.CategoryID = Tasks[j].CategoryID
				tempTask.CreatedAt = Tasks[j].CreatedAt
				tempTask.UpdatedAt = Tasks[j].UpdatedAt
				tempGetCategory.Tasks = append(tempGetCategory.Tasks, tempTask)
			}

		}
		getCategory = append(getCategory, tempGetCategory)

	}
	c.JSON(http.StatusOK, getCategory)
}

func CategorieEdit(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	//Category := model.Category{}
	newCategory := model.Category{}

	categoryID, _ := strconv.Atoi(c.Param("categoryId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&newCategory)
	} else {
		c.ShouldBind(&newCategory)
	}

	newCategory.ID = uint(categoryID)

	db.Model(&newCategory).Where("id = ?", categoryID).Update("type", newCategory.Type)

	c.JSON(http.StatusOK, gin.H{
		"id":         newCategory.ID,
		"type":       newCategory.Type,
		"updated_at": newCategory.UpdatedAt,
	})
}

func CategorieDelete(c *gin.Context) {
	db := database.GetDB()
	categoryID, _ := strconv.Atoi(c.Param("categoryId"))
	Category := model.Category{}
	db.Delete(&Category, categoryID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Category has been successfully deleted",
	})

}
