package http

import (
	"log"

	"github.com/lalo64/SmartEnv-api/src/users/application"
	"github.com/lalo64/SmartEnv-api/src/users/application/services"
	"github.com/lalo64/SmartEnv-api/src/users/domain/ports"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/adapters"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http/controllers"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http/controllers/helper"
)

var (
    userRepository  ports.IUserRepository
    bcryptService   services.BcryptService 
)

func init() {
    var err error

    // Inicializar el repositorio de usuarios
    userRepository, err = adapters.NewUserRepository()
    if err != nil {
        log.Fatalf("Error initializing user repository: %v", err)
    }

    // Inicializar el servicio de encriptación
    bcryptService, err = helper.NewBcryptHelper()
    if err != nil {
        log.Fatalf("Error initializing bcrypt service: %v", err)
    }
}

func SetUpRegisterController() *controllers.CreateUserController {
    createUserUseCase := application.NewCreateUserUseCase(userRepository, bcryptService)
    return controllers.NewCreateUserController(createUserUseCase)
}

func GetUserByIDController() *controllers.GetUserByIDController {
    getUserByIDUseCase := application.NewUserGetByIDUseCase(userRepository)
    return controllers.NewGetUserByIDUseCase(getUserByIDUseCase)
}

func DeleteUserController() *controllers.DeleteUserController {
    deleteUserUseCase := application.NewDeleteUserUseCase(userRepository)
    return controllers.NewDeleteUserController(deleteUserUseCase)
}

func CheckEmailController() *controllers.CheckEmailController {
    checkEmailUseCase := application.NewCheckEmailUseCase(userRepository)
    return controllers.NewCheckEmailController(checkEmailUseCase)
}

func AuthController() *controllers.AuthController {
	authUseCase := application.NewAuthUseCase(userRepository)
    return controllers.NewAuthController(authUseCase)
}