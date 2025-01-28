package application

import (
	"github.com/lalo64/SmartEnv-api/src/users/domain/ports"
)

type CheckEmailUseCase struct {
	UserRepository ports.IUserRepository
}


func NewCheckEmailUseCase(userRepository ports.IUserRepository) *CheckEmailUseCase {
	return &CheckEmailUseCase{UserRepository: userRepository}
}

func (s CheckEmailUseCase) Run(email string) (bool, error) {
	status, err := s.UserRepository.CheckEmail(email)

	if err != nil {
		return false, err
	}

	return status, nil
}