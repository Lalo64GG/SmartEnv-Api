package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http"
)

func UserRoutes(router *gin.RouterGroup){
	createUserController := http.SetUpRegisterController()
	getUserByIDController := http.GetUserByIDController()
	deleteUserController := http.DeleteUserController()

	router.POST("/", createUserController.Run)
	router.GET("/:id", getUserByIDController.Run)
	router.DELETE("/:id", deleteUserController.Run)
}