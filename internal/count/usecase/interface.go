package usecase

type Provider interface {
	SelectQuery() (int, error)
	InsertQuery(int) (bool, error)
	SetQuery(int) (bool, error)
	ClearQuery() (bool, error)
}
