package usecase

import (
	"errors"
	"fmt"
	"strconv"
)

// Возвращает текущий счётчик
func (u *Usecase) SelectCounter() (string, error) {
	count, err := u.p.SelectQuery()
	if err != nil {
		// -1 возвращается, если ошибка, НО не из-за пустой таблицы
		if count == -1 {
			return u.errorMsg, err
		}
	}
	return fmt.Sprintf("Счётчик сейчас %d", count), nil
}

// Увеличивает счётчик на 1
func (u *Usecase) PostCounter() (string, error) {
	count, _ := u.p.SelectQuery()
	count += 1
	var err error
	var done bool
	if count > 1 {
		done, err = u.p.SetQuery(count)
	} else {
		done, err = u.p.InsertQuery(count)
	}
	if done {
		return u.defaultMsg, nil
	}
	return u.errorMsg, err
}

// Устанавливает заданное значение счётчика
func (u *Usecase) SetCounter(msg string) (string, error) {
	var count int
	var err error
	if msg == "" {
		count, _ = u.p.SelectQuery()
	} else {
		count, err = strconv.Atoi(msg)
	}
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			return u.inputErrorMsg, nil
		}
		return u.errorMsg, err
	}
	_, err = u.p.SetQuery(count)
	if err != nil {
		return u.errorMsg, err
	}
	return fmt.Sprintf("Значение %d установлено", count), nil
}

// Обнуляет счётчик
func (u *Usecase) ClearCounter() (string, error) {
	_, err := u.p.ClearQuery()
	if err != nil {
		return u.errorMsg, err
	}
	return "Счетик сброшен...", nil
}
