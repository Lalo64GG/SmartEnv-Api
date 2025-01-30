package application

import (
	"github.com/lalo64/SmartEnv-api/src/users/domain/entities"
	"github.com/lalo64/SmartEnv-api/src/users/domain/ports"
)

type AuthUseCase struct {
	UserRepository ports.IUserRepository
}

func NewAuthUseCase(userRepository ports.IUserRepository) *AuthUseCase {
	return &AuthUseCase{UserRepository: userRepository}  
}


func (s AuthUseCase) Run(email string)(entities.User, error){
	user, err := s.UserRepository.GetUserByEmail(email)

	if err != nil {
        return entities.User{}, err
    }

	return user, nil
}