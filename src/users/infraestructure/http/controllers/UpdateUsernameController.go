package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/SmartEnv-api/src/shared/responses"
	"github.com/lalo64/SmartEnv-api/src/users/application"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http/request"
)

type UpdateController struct {
	UpdateUseCase *application.UpdateUsernameUseCase
}

func NewUpdateController(updateUseCase *application.UpdateUsernameUseCase) *UpdateController {
	return &UpdateController{UpdateUseCase: updateUseCase}
}

func (ctr UpdateController) Run(ctx *gin.Context) {
	var req request.UpdateUserRequest


	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid ID",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	user, err := ctr.UpdateUseCase.Run(id, req.Username)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Error updating username",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, responses.Response{
		Success: true,
		Message: "Username updated",
		Data:    user.Username,
		Error:   nil,
	})
}
