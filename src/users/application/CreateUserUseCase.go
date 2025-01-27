package application

import (
	"github.com/lalo64/SmartEnv-api/src/users/domain/entities"
	"github.com/lalo64/SmartEnv-api/src/users/domain/ports"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	userRepository ports.IUserRepository
}


func NewCreateUserUseCase(userRepository ports.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepository: userRepository}
}

func (s *CreateUserUseCase) Run (Username, Email, password string) (entities.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return entities.User{}, err
	}

	userCre := entities.User{
		Username: Username,
		Email: Email,
		Password: string(hashedPassword),
	}

	newUser, err := s.userRepository.Create(userCre)

	if err != nil {
		return entities.User{}, err
	}

	return  newUser, nil
}
