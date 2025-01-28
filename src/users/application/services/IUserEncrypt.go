package services

import()

type IUserEncrypt struct {
	Encrypt func(password string) (string, error)
}  