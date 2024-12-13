package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) SelectQuery() (int, error) {
	var c int
	err := p.conn.QueryRow("SELECT c FROM counter LIMIT 1").Scan(&c)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return -1, err
	}

	return c, nil
}

func (p *Provider) InsertQuery(c int) (bool, error) {
	_, err := p.conn.Exec("INSERT INTO counter (c) VALUES ($1)", c)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Provider) SetQuery(c int) (bool, error) {
	_, err := p.conn.Exec("UPDATE counter SET c=$1", c)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Provider) ClearQuery() (bool, error) {
	_, err := p.conn.Exec("DELETE FROM counter")

	if err != nil {
		return false, err
	}

	return true, nil
}
