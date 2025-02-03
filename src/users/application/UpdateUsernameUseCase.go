package application

import (
	"github.com/lalo64/SmartEnv-api/src/users/domain/entities"
	"github.com/lalo64/SmartEnv-api/src/users/domain/ports"
)

type UpdateUsernameUseCase struct {
	UpdateService ports.IUserRepository
}

func NewUpdateUserUseCase(updateService ports.IUserRepository) *UpdateUsernameUseCase {
	return &UpdateUsernameUseCase{UpdateService: updateService}
}


func (s *UpdateUsernameUseCase) Run(id int, username string) (entities.User, error){
	user := entities.User{
		ID: id,
		Username: username,
	}

	user, err := s.UpdateService.UpdateUsername(user)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}