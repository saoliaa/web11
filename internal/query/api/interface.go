package api

type Usecase interface {
	GetQuery() (string, error)
	PostQuery(string, string) (string, error)
	ClearQuery() (string, error)
}
