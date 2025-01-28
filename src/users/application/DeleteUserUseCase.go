package application

import "github.com/lalo64/SmartEnv-api/src/users/domain/ports"

type DeleteUserUseCase struct {
	UserRepository ports.IUserRepository
}

func NewDeleteUserUseCase(userRepository ports.IUserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{ UserRepository: userRepository }
}


func (s *DeleteUserUseCase) Run(id int64) (bool, error){
	_, err := s.UserRepository.DeleteUser(id)

	if err != nil {
		return false, err  
	}
	
	return true, nil 
}