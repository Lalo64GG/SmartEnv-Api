package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http"
)

func UserRoutes(router *gin.RouterGroup){
	createUserController := http.SetUpRegisterController()

	router.POST("/", createUserController.Run)
}