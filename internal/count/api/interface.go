package api

type Usecase interface {
	SelectCounter() (string, error)
	PostCounter() (string, error)
	SetCounter(string) (string, error)
	ClearCounter() (string, error)
}
