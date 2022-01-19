package controller

import (
	"fmt"
	"my_kanban_board/database"
	"my_kanban_board/helper"
	"my_kanban_board/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType

	User := model.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	password = User.Password
	err := db.Debug().Where("email=?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}
	comparePass := helper.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helper.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	User := model.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.Role = "member"

	err := db.Debug().Create(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         User.ID,
		"full_name":  User.FullName,
		"email":      User.Email,
		"created_at": User.CreatedAt,
	})

}

func UserUpdate(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	newUser := model.User{}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	//userID, _ := strconv.Atoi(c.Param("user_id"))
	User := model.User{}
	_ = db.First(&User, userID)

	if contentType == appJSON {
		c.ShouldBindJSON(&newUser)
	} else {
		c.ShouldBind(&newUser)
	}

	newUser.Password = helper.HashPass(newUser.Password)
	oldUser := model.User{}
	db.Where("id = ?", userID).Find(&oldUser)
	fmt.Println(oldUser)

	newUser.ID = uint(userID)
	newUser.Password = oldUser.Password
	newUser.Role = oldUser.Role

	err := db.Model(&newUser).Where("id=?", userID).Updates(
		model.User{
			FullName: newUser.FullName,
			Email:    newUser.Email,
			Password: newUser.Password,
			Role:     newUser.Role,
		}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         newUser.ID,
		"full_name":  newUser.FullName,
		"email":      newUser.Email,
		"updated_at": newUser.UpdatedAt,
	})
}

func UserDelete(c *gin.Context) {
	db := database.GetDB()
	User := model.User{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	db.Delete(&User, userID)

	c.JSON(http.StatusOK, gin.H{
		"message": "User has been deleted",
	})
}
