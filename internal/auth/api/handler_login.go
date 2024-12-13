package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

// Структура HTTP-запроса на вход в аккаунт
type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Структура HTTP-ответа на вход в аккаунт
// В ответе содержится JWT-токен авторизованного пользователя
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

var (
	errBadCredentials = errors.New("email or password is incorrect")
)

// Секретный ключ для подписи JWT-токена
// Необходимо хранить в безопасном месте
var jwtSecretKey = []byte("very-secret-key")

// Обработчик HTTP-запросов на вход в аккаунт
// Логика входа в систему (Login)
func (srv *Server) Login(c *fiber.Ctx) error {
	loginReq := LoginRequest{}
	if err := c.BodyParser(&loginReq); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	_, login := srv.uc.Login(loginReq.Login, loginReq.Password)
	if !login {
		return errBadCredentials
	}

	payload := jwt.MapClaims{
		"sub": loginReq.Login,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString(jwtSecretKey) // Используем тот же ключ
	if err != nil {
		logrus.WithError(err).Error("JWT token signing")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(LoginResponse{AccessToken: t})
}
