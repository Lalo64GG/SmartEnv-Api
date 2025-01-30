package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/SmartEnv-api/src/shared/middlewares"
	"github.com/lalo64/SmartEnv-api/src/shared/responses"
	"github.com/lalo64/SmartEnv-api/src/users/application"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http/controllers/helper"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http/request"
)

type AuthController struct {
	AuthService  *application.AuthUseCase
	BcryptHelper *helper.BcryptHelper
}


func NewAuthController(authService *application.AuthUseCase) *AuthController {
	return &AuthController{AuthService: authService}
}


func (ctr *AuthController) Run(ctx *gin.Context) {

	var authRequest request.AuthRequest

	if err := ctx.ShouldBind(&authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Llene todos los campos",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	user, err := ctr.AuthService.Run(authRequest.Email)

	if err != nil {
		switch err.Error() {
		case "sql: no rows in result set":
			ctx.JSON(http.StatusNotFound, responses.Response{
				Success: false,
				Message: "Usuario no encontrado",
				Data:    nil,
				Error:   err.Error(),
			})
			default:
				ctx.JSON(http.StatusInternalServerError, responses.Response{
					Success: false,
					Message: "Error al inciar sesión",
					Data:    nil,
					Error:   err.Error(),
				})
			}
			return 
	}

	if err := ctr.BcryptHelper.Compare(user.Password, []byte(authRequest.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, responses.Response{
			Success: false,
			Message: "Contraseña incorrecta",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	token, err := middlewares.GenerateJWT(int64(user.ID), user.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
            Success: false,
            Message: "Error generando el token",
            Data:    nil,
            Error:   err.Error(),
        })
        return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Sesión iniciada",
		Data:   map[string]interface{}{
			"Token": token,
		},
		Error:   nil,
	})


}