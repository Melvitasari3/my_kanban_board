package middleware

import (
	"my_kanban_board/database"
	"my_kanban_board/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// db := database.GetDB()
		// // userID, err := strconv.Atoi(c.Param("user_id"))
		// // if err != nil {
		// // 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		// // 		"error":   "Bad request",
		// // 		"message": "invalid parameter",
		// // 	})
		// // 	return
		// // }
		// userData := c.MustGet("userData").(jwt.MapClaims)
		// userID := uint(userData["id"].(float64))
		// User := model.User{}
		// err := db.Select("id").First(&User, uint(userID)).Error

		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		// 		"error":   "Unauthorized",
		// 		"message": "you are not allow to access this data",
		// 	})
		// 	return
		// }
		// // if User.ID != userID {
		// // 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		// // 		"error":   "Unauthorized",
		// // 		"message": "you are not allow to access this data",
		// // 	})
		// // 	return
		// // }

		c.Next()
	}
}

func CategoryAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := model.User{}

		err := db.Debug().Where("id=?", userID).Take(&User).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if User.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func TaskAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		taskID, _ := strconv.Atoi(c.Param("taskId"))
		Task := model.Task{}
		err := db.Where("id = ?", taskID).First(&Task)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "data not found",
				"message": "data doesn't exist",
			})
			return
		}
		if userID != Task.UserID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allow to access this data",
			})
			return
		}

	}
}
