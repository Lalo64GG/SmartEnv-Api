package adapters

import (
	"database/sql"
	"log"

	"github.com/lalo64/SmartEnv-api/src/config"
	"github.com/lalo64/SmartEnv-api/src/users/domain/entities"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository() (*UserRepository, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}

	return &UserRepository{DB: db}, nil
}

func (r *UserRepository) Create(user entities.User) (entities.User, error) {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`

	stmt, err := r.DB.Prepare(query)

	if err != nil {
		log.Fatal(err, 1)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Email, user.Password)

	if err != nil {
		return entities.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entities.User{}, err
	}

	user.ID = int(id)
	user.Password = ""

	return user, nil
}

func (r *UserRepository) GetByID(id int64) (entities.User, error) {
	query := `SELECT id, username, email FROM users WHERE id = ?`

	stmt, err := r.DB.Prepare(query)

	if err != nil {
		log.Fatal(err, 1)
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var user entities.User

	err = row.Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (r *UserRepository) CheckEmail(email string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = ?) AS existe`

	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	var existe bool
	log.Print(existe, email)
	err = stmt.QueryRow(email).Scan(&existe)

	if err != nil {
		log.Print(err, 2)
		return false, err
	}

	return existe, nil
}

func (r *UserRepository) DeleteUser(id int64) (bool, error) {
	query := `DELETE FROM users WHERE id = ?`

	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepository) GetUserByEmail(email string) (entities.User, error) {

	query := `SELECT id, email, password, username FROM users WHERE email = ?`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return entities.User{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)

	var user entities.User

	err = row.Scan(&user.ID, &user.Email, &user.Password, &user.Username)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil

}

func (r *UserRepository) UpdateUsername(user entities.User) (entities.User, error) {
	query := `UPDATE users SET username = ? WHERE id = ?`

	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return entities.User{}, err
	}

	defer stmt.Close()

	// Corregir: Pasar directamente los valores de user.ID y user.Username
	_, err = stmt.Exec(user.Username, user.ID)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
