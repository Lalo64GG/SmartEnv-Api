package helper

import "golang.org/x/crypto/bcrypt"

func Encrypt(pwd []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}