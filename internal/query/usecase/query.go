package usecase

import (
	"fmt"
	"strconv"
)

// Возвращает случайное имя и возраст
func (u *Usecase) GetQuery() (string, error) {
	name, age, err := u.p.SelectQuery()
	if err != nil {
		if age == -1 {
			return u.errorMsg, err
		}
	}
	return fmt.Sprintf("Name=%s Age=%d", name, age), nil
}

func (u *Usecase) PostQuery(name string, age string) (string, error) {
	if name == "" {
		name = "User"
	}
	if age == "" {
		age = "0"
	}
	int_age, err := strconv.Atoi(age)
	if err != nil {
		return u.inputErrorMsg, err
	}
	_, err = u.p.InsertQuery(name, int_age)
	if err != nil {
		return u.errorMsg, err
	}
	return u.defaultMsg, err
}

func (u *Usecase) ClearQuery() (string, error) {
	_, err := u.p.ClearQuery()
	if err != nil {
		return u.errorMsg, err
	}
	return "Значения сброшены...", nil
}
