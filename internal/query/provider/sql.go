package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) SelectQuery() (string, int, error) {
	var name string
	var age int
	err := p.conn.QueryRow("SELECT name, age FROM query ORDER BY RANDOM() LIMIT 1").Scan(&name, &age)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "User", 0, nil
		}
		return "Error", -1, err
	}

	return name, age, nil
}

func (p *Provider) InsertQuery(name string, age int) (bool, error) {
	_, err := p.conn.Exec("INSERT INTO query (name, age) VALUES ($1, $2)", name, age)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Provider) ClearQuery() (bool, error) {
	_, err := p.conn.Exec("DELETE FROM query")

	if err != nil {
		return false, err
	}

	return true, nil
}
