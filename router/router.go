package router

import (
	"my_kanban_board/controller"
	"my_kanban_board/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/login", controller.UserLogin)
		userRouter.POST("/register", controller.UserRegister)

		userRouter.Use(middleware.Authentication())
		userRouter.PUT("/update-account", controller.UserUpdate)
		userRouter.DELETE("/delete-account", controller.UserDelete)
	}

	categorieRouter := r.Group("/categories")
	{
		categorieRouter.Use(middleware.Authentication())
		categorieRouter.POST("/", middleware.CategoryAuthorization(), controller.CategoriePost)
		categorieRouter.GET("/", controller.CategorieViewAll)
		categorieRouter.PATCH("/:categoryId", middleware.CategoryAuthorization(), controller.CategorieEdit)
		categorieRouter.DELETE("/:categoryId", middleware.CategoryAuthorization(), controller.CategorieDelete)
	}

	taskRouter := r.Group("/tasks")
	{
		taskRouter.Use(middleware.Authentication())
		taskRouter.POST("/", controller.TaskPost)
		taskRouter.GET("/", controller.TaskView)
		taskRouter.PUT("/:taslId", middleware.TaskAuthorization(), controller.CategorieEdit)
		taskRouter.PATCH("/update-status/:taskId", middleware.TaskAuthorization(), controller.TaskUpdateStatus)
		taskRouter.PATCH("/update-category/:taskId", middleware.TaskAuthorization(), controller.TaskUpdateCategory)
		taskRouter.DELETE("/:taskId", middleware.TaskAuthorization(), controller.TaskDelete)
	}

	return r
}
