package ports

import "github.com/lalo64/SmartEnv-api/src/users/domain/entities"

type IUserRepository interface {
	Create(user entities.User) (entities.User, error)
	GetByID(id int64) (entities.User, error)
	CheckEmail(email string) (bool, error)
	DeleteUser(id int64) (bool, error)
}