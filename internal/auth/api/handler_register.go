package api

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (srv *Server) Register(c *fiber.Ctx) error {
	regReq := RegisterRequest{}
	if err := c.BodyParser(&regReq); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	// Проверяем, что пользователь с таким email еще не зарегистрирован
	if _, exists := srv.uc.Exist(regReq.Login); exists {
		return errors.New("the user already exists")
	}

	srv.uc.Register(
		regReq.Login,
		regReq.Password,
		regReq.Name)

	return c.SendStatus(fiber.StatusCreated)
}
