package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/SmartEnv-api/src/records/application"
	"github.com/lalo64/SmartEnv-api/src/shared/responses"
)

type GetAllRecordsController struct {
	GetAllRecordsController *application.GetAllRecordsUseCase
}

func NewGetAllRecordsController(getAllRecordsService *application.GetAllRecordsUseCase) *GetAllRecordsController{
	return &GetAllRecordsController{GetAllRecordsController: getAllRecordsService}
}


func (ctr *GetAllRecordsController) Run(ctx *gin.Context){
	limit := parseQueryParam(ctx, "limit", 5)
	page := parseQueryParam(ctx, "page", 1)
	orderBy := ctx.DefaultQuery("orderBy", "id")
	orderDir := ctx.DefaultQuery("orderDir", "asc")

	if orderDir != "asc" && orderDir != "desc" {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid value for orderDir, must be 'asc' or 'desc'",
		})
		return
	}

	records, err := ctr.GetAllRecordsController.Run(limit, page, orderBy, orderDir)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
            Success: false,
            Message: "Error obteniendo los registros",
            Data:    nil,
            Error:   err.Error(),
        })
        return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
        Message: "Registros obtenidos con exito",
        Data:    records,
        Error:   nil,
	})

}

func parseQueryParam(ctx *gin.Context, key string, defaultValue int64) int64 {
	queryValue := ctx.Query(key)
	if queryValue == "" {
		return defaultValue
	}
	value, err := strconv.ParseInt(queryValue, 10, 64)
	if err != nil || value <= 0 {
		return defaultValue
	}
	return value
}
