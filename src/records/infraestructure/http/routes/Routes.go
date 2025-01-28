package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/SmartEnv-api/src/records/infraestructure/http"
)

func RecordRoutes(router *gin.RouterGroup){
	createRecord := http.SetUpRegisterController()
	getRecordByID := http.GetRecordByIDController()
	getAllRecord := http.GetAllRecordsController()

	router.POST("/", createRecord.Run)
	router.GET("/:id", getRecordByID.Run)
	router.GET("/", getAllRecord.Run)
}