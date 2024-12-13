package usecase

// Возвращает текущий счётчик
func (u *Usecase) Register(login string, password string, name string) (string, bool) {
	_, err := u.p.SelectLogin(login)
	if err == nil {
		return "Пользователь существует", false
	}
	u.p.CreateUser(login, password, name)
	return "Пользователь создан", true
}

// Увеличивает счётчик на 1
func (u *Usecase) Login(login string, password string) (string, bool) {
	sample_password, err := u.p.CheckPassword(login)
	if err != nil {
		return "Пользователь не найден", false
	}
	if password != sample_password {
		return "Неверный пароль", false
	}
	name, _ := u.p.SelectLogin(login)
	return name, true
}

// Устанавливает заданное значение счётчика
func (u *Usecase) Exist(login string) (string, bool) {
	name, err := u.p.SelectLogin(login)
	if err != nil {
		return "Пользователь не найден", false
	}
	return name, true
}
