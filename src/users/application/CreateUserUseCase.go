package application

import (
	"github.com/lalo64/SmartEnv-api/src/users/application/services"
	"github.com/lalo64/SmartEnv-api/src/users/domain/entities"
	"github.com/lalo64/SmartEnv-api/src/users/domain/ports"
)

type CreateUserUseCase struct {
	userRepository ports.IUserRepository
	BycryptService services.BcryptService
}


func NewCreateUserUseCase(userRepository ports.IUserRepository, bycryptService services.BcryptService) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: userRepository, BycryptService: bycryptService}
}

func (s *CreateUserUseCase) Run(Username, Email, password string) (entities.User, error) {

	encryptedPass, err := s.BycryptService.Encrypt([]byte(password))

	if err != nil {
		return entities.User{}, err    
	}
	


	userCre := entities.User{
		Username: Username,
		Email: Email,
		Password: encryptedPass,
	}

	newUser, err := s.userRepository.Create(userCre)

	if err != nil {
		return entities.User{}, err
	}

	return  newUser, nil
}
