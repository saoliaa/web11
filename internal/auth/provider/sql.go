package provider

func (p *Provider) SelectLogin(login string) (string, error) {
	var name string
	err := p.conn.QueryRow("SELECT name FROM users WHERE login=$1", login).Scan(&name)

	if err != nil {
		return "None", err
	}

	return name, nil
}

func (p *Provider) CreateUser(login string, password string, name string) (bool, error) {
	_, err := p.conn.Exec("INSERT INTO users (login, password, name) VALUES ($1, $2, $3)", login, password, name)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Provider) CheckPassword(login string) (string, error) {
	var password string
	err := p.conn.QueryRow("SELECT password FROM users WHERE login=$1", login).Scan(&password)

	if err != nil {
		return "None", err
	}

	return password, nil
}
