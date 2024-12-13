package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Читаем тело запроса
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Восстанавливаем тело запроса для следующего обработчика
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// Получаем токен из заголовков
		token := c.Request().Header.Get("Authorization")

		// Создаем новый запрос для отправки на сервис localhost:8021/auth
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8021/auth", bytes.NewBuffer(body))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		req.Header.Set("Authorization", token)

		// Отправляем запрос и получаем ответ
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		defer resp.Body.Close()

		// Проверяем статус ответа
		logrus.Error("status: ", resp.Body)
		if resp.StatusCode != http.StatusOK {
			return echo.NewHTTPError(resp.StatusCode, "Unauthorized")
		}

		// Если всё в порядке, передаем управление следующему обработчику
		return next(c)
	}
}
