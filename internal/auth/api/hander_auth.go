package api

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

// Структура HTTP-ответа с информацией о пользователе
type Response struct {
	Name string `json:"name"`
}

// Обработка токена (jwtPayloadFromRequest)
func jwtPayloadFromRequest(c *fiber.Ctx) (jwt.MapClaims, bool) {
	header := c.Get("Authorization")
	if header == "" {
		logrus.Error("Authorization header is missing")
		return nil, false
	}

	tokenString := strings.TrimPrefix(header, "Bearer ")
	if tokenString == header {
		logrus.Error("Invalid token format")
		return nil, false
	}

	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil // Используем тот же ключ
	})

	if err != nil || !jwtToken.Valid {
		logrus.Error("Invalid JWT token: ", err)
		return nil, false
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"jwt_token_claims": jwtToken.Claims,
		}).Error("wrong type of JWT token claims")
		return nil, false
	}

	return payload, true
}

// Обработчик HTTP-запросов на получение информации о пользователе
func (srv *Server) Auth(c *fiber.Ctx) error {
	jwtPayload, ok := jwtPayloadFromRequest(c)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	name, ok := srv.uc.Exist(jwtPayload["sub"].(string))
	if !ok {
		return errors.New("user not found")
	}

	return c.JSON(Response{Name: name})
}
