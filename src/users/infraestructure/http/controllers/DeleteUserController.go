package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/SmartEnv-api/src/shared/responses"
	"github.com/lalo64/SmartEnv-api/src/users/application"
)

type DeleteUserController struct {
	DeleteUserController *application.DeleteUserUseCase
}

func NewDeleteUserController(deleteUserService *application.DeleteUserUseCase) *DeleteUserController{
	return &DeleteUserController{DeleteUserController: deleteUserService}
}


func (ctr *DeleteUserController) Run(ctx *gin.Context) {
	idParams := ctx.Param("id")

	id, err := strconv.ParseInt(idParams, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
            Message: "ID inválido",
            Data:    nil,
            Error:   err.Error(),
		})
		return
	}


	_, error := ctr.DeleteUserController.Run(id)

	if error != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
            Success: false,
            Message: "Error al eliminar el usuario",
            Data:    nil,
            Error:   error.Error(),
        })
        return
	}


	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
        Message: "Usuario eliminado correctamente",
        Data:    nil,
        Error:   "",
	})

}