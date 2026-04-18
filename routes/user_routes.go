package routes

import (
	"belajar-go/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userController *controllers.UserController) {
	users := r.Group("/users")
	{
		users.POST("", userController.CreateUser)
		users.GET("", userController.GetUsers)
		users.DELETE("/:id", userController.DeleteUser)
		users.PUT("/:id", userController.UpdateUser)
	}
}
