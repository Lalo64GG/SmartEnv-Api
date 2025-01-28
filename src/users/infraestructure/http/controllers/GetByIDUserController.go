package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/SmartEnv-api/src/shared/responses"
	"github.com/lalo64/SmartEnv-api/src/users/application"
)


type GetUserByIDController struct {
	GetByIdUserController *application.GetUserByIDUseCase
}

func NewGetUserByIDUseCase(getByIDUserService *application.GetUserByIDUseCase) *GetUserByIDController {
	return &GetUserByIDController{GetByIdUserController: getByIDUserService}
}

func (ctr *GetUserByIDController) Run(ctx *gin.Context){
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err!= nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
            Success: false,
            Message: "ID invalido",
            Data:    nil,
            Error:   err.Error(),
        })
        return     }

	 user, err := ctr.GetByIdUserController.Run(id)

	 if err != nil {
        ctx.JSON(http.StatusInternalServerError, responses.Response{
            Success: false,
            Message: "Error obteniendo el usuario",
            Data:    nil,
            Error:   err.Error(),
        })
        return     }

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
        Message: "Usuario obtenido con exito",
        Data:    user,
        Error:   nil,
	})
}