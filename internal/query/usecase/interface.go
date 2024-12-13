package usecase

type Provider interface {
	SelectQuery() (string, int, error)
	InsertQuery(string, int) (bool, error)
	ClearQuery() (bool, error)
}
