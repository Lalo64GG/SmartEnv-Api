package application

import (
	"github.com/lalo64/SmartEnv-api/src/users/domain/entities"
	"github.com/lalo64/SmartEnv-api/src/users/domain/ports"
)

type GetUserByIDUseCase struct {
	UserRepository ports.IUserRepository
}

func NewUserGetByIDUseCase( userRepository ports.IUserRepository) *GetUserByIDUseCase{
	return &GetUserByIDUseCase{UserRepository: userRepository} 
}

func (s *GetUserByIDUseCase) Run(id int64) (entities.User, error) {

	user, err := s.UserRepository.GetByID(id)

	if err != nil {
        return entities.User{}, err
    }

	return user, nil
}