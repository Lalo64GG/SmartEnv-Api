package http

import (
	"log"

	"github.com/lalo64/SmartEnv-api/src/users/application"
	"github.com/lalo64/SmartEnv-api/src/users/domain/ports"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/adapters"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http/controllers"
)

var userRepository ports.IUserRepository


func init(){
	var err error
	userRepository, err = adapters.NewUserRepository()
	if err != nil {
		log.Fatalf("Error initializing user repository: %v", err)
	}
}


func SetUpRegisterController() *controllers.CreateUserController{
	createService := application.NewCreateUserUseCase(userRepository)
	return controllers.NewCreateUserController(createService)
}