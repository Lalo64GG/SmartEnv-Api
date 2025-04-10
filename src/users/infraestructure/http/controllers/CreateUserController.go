package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lalo64/SmartEnv-api/src/shared/responses"
	"github.com/lalo64/SmartEnv-api/src/users/application"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http/controllers/helpers"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http/request"
)

type CreateUserController struct {
	CreateUserController *application.CreateUserUseCase
	Validator            *validator.Validate
	BcryptHelper 			 *helper.BcryptHelper
}

func NewCreateUserController(createUserService *application.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{CreateUserController: createUserService, Validator: validator.New()}
}


func (ctr *CreateUserController) Run(ctx *gin.Context) {
	var req request.CreateUserRequest

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


		user, err := ctr.CreateUserController.Run(req.Username, req.Email, req.Password)

		if err != nil {

			if strings.Contains(err.Error(), "unique_user_email") {
				ctx.JSON(http.StatusBadRequest, responses.Response{
					Success: false,
					Message: "El correo ingresado ya se encuentra registrado",
					Data:    nil,
					Error:   err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusBadRequest, responses.Response{
				Success: false,
				Message: err.Error(),
				Error:   err.Error(),
			})
			return
		}

		


	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Usuario creado correctamente",
		Data:    user,
        Error:   "",
	})
}