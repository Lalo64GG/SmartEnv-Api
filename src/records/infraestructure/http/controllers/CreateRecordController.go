package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lalo64/SmartEnv-api/src/records/application"
	"github.com/lalo64/SmartEnv-api/src/records/infraestructure/http/request"
	"github.com/lalo64/SmartEnv-api/src/shared/responses"
)

type CreateRecordController struct {
	RecordController *application.CreateRecordUseCase
	Validator 	  *validator.Validate
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

	record, err := ctr.RecordController.Run(req.Temperature, req.Distance)

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