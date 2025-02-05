package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lalo64/SmartEnv-api/src/records/application"
	"github.com/lalo64/SmartEnv-api/src/records/domain/entities"
	"github.com/lalo64/SmartEnv-api/src/records/infraestructure/http/controllers/helpers"
	"github.com/lalo64/SmartEnv-api/src/records/infraestructure/http/request"
	"github.com/lalo64/SmartEnv-api/src/shared/responses"
)

type CreateRecordController struct {
	RecordController *application.CreateRecordUseCase
	Validator 	  	 *validator.Validate
	KafkaService   	 *helpers.KafkaHelper
}

func NewCreateRecordController (recordService *application.CreateRecordUseCase) *CreateRecordController{
	return &CreateRecordController{RecordController: recordService, Validator: validator.New()}
}

func (ctr *CreateRecordController) Run(ctx *gin.Context){

	var req request.CreateRecordRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
            Message: "Llene todos los campos",
            Data:    nil,
            Error:   err.Error(),
		})
	}

	if err := ctr.Validator.Struct(req); err != nil {
		validationErrors  := err.(validator.ValidationErrors)
        ctx.JSON(http.StatusBadRequest, responses.Response{
                Success: false,
                Message: "Errores de validación",
                Data:    nil,
                Error:   validationErrors.Error(),
		})
		return
	}

	record := entities.Record{
		Temperature: req.Temperature,
		Humidity: req.Humidity,
		Gas_level: req.Gas_level,

	}

	success, error := ctr.KafkaService.Producer(record)

	if error!= nil || !success {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error enviando el mensaje",
			Data:    nil,
			Error:   error.Error(),
		})
	}


	record, err := ctr.RecordController.Run(req.Temperature, req.Humidity, req.Gas_level)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
            Success: false,
            Message: "Error creando el registro",
            Data:    nil,
            Error:   err.Error(),
        })
        return     
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
        Message: "Registro creado correctamente",
        Data:    record,
        Error:   "",
	})

}