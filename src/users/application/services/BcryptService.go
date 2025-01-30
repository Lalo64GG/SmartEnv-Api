package services

type BcryptService interface {
	Encrypt(pwd []byte) (string, error)
	Compare(hashedPwd string, plainPwd []byte)  error
}
