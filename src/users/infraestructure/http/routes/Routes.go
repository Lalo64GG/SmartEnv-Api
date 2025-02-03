package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http"
)

func UserRoutes(router *gin.RouterGroup){
	createUserController := http.SetUpRegisterController()
	getUserByIDController := http.GetUserByIDController()
	deleteUserController := http.DeleteUserController()
	checkEmailController := http.CheckEmailController()
	authController := http.AuthController()
	updateController := http.SetUpUpdateUserController()

	router.POST("/", createUserController.Run)
	router.POST("/login", authController.Run)
	router.GET("/:id", getUserByIDController.Run)
	router.GET("/verificate/:email", checkEmailController.Run)
	router.DELETE("/:id", deleteUserController.Run)
	router.PUT("/:id", updateController.Run)
}