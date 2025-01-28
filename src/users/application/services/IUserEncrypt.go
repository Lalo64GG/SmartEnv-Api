package services

type IUserEncrypt interface {
	Encrypt(pwd []byte) (string, error)
}  