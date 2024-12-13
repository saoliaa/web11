package api

type Usecase interface {
	Register(login string, password string, name string) (string, bool)
	Login(login string, password string) (string, bool)
	Exist(login string) (string, bool)
}
