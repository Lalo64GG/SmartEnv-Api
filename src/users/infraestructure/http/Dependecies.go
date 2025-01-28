package http

import (
	"log"

	"github.com/lalo64/SmartEnv-api/src/users/application"
	"github.com/lalo64/SmartEnv-api/src/users/application/services"
	"github.com/lalo64/SmartEnv-api/src/users/domain/ports"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/adapters"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http/controllers"
)

var userRepository ports.IUserRepository
var IUserEncrypt  services.IUserEncrypt


func init(){
	var err error
	userRepository, err = adapters.NewUserRepository()
	if err != nil {
		log.Fatalf("Error initializing user repository: %v", err)
	}
}


func SetUpRegisterController() *controllers.CreateUserController{
	createService := application.NewCreateUserUseCase(userRepository, IUserEncrypt)
	return controllers.NewCreateUserController(createService)
}

func GetUserByIDController() *controllers.GetUserByIDController{
	getByIDService := application.NewUserGetByIDUseCase(userRepository)
	return controllers.NewGetUserByIDUseCase(getByIDService)
}

func DeleteUserController() *controllers.DeleteUserController {
	deleteUserService := application.NewDeleteUserUseCase(userRepository)
	return controllers.NewDeleteUserController(deleteUserService)
}

func CheckEmailController() *controllers.CheckEmailController {
	checkEmailService := application.NewCheckEmailUseCase(userRepository)
    return controllers.NewCheckEmailController(checkEmailService)
}