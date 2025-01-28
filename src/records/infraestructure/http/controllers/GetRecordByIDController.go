package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/SmartEnv-api/src/records/application"
	"github.com/lalo64/SmartEnv-api/src/shared/responses"
)

type GetByIdRecordController struct {
	GetByIdRecordController *application.GetRecordByIDUseCase
}

func NewGetByIdRecordController(getByIdService *application.GetRecordByIDUseCase) *GetByIdRecordController{
	return &GetByIdRecordController{GetByIdRecordController: getByIdService}  
}

func(ctr *GetByIdRecordController) Run(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
            Message: "ID inválido",
            Data:    nil,
            Error:   err.Error(),
		})
	}

	record, err := ctr.GetByIdRecordController.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
            Success: false,
            Message: "Error obteniendo el registro",
            Data:    nil,
            Error:   err.Error(),
        })
        return     
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
        Message: "Registro obtenido con éxito",
        Data:    record,
        Error:   nil,
	})
	
}