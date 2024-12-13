package usecase

type Provider interface {
	SelectLogin(login string) (string, error)
	CreateUser(login string, password string, name string) (bool, error)
	CheckPassword(login string) (string, error)
}
