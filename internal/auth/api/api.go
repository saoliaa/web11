package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Server struct {
	maxSize int

	server  *fiber.App
	address string

	uc Usecase
}

func NewServer(ip string, port int, maxSize int, uc Usecase) *Server {
	api := Server{
		maxSize: maxSize,
		uc:      uc,
	}

	api.server = fiber.New()
	// api.server.Use(logger.New())

	api.server.Post("/register", api.Register)
	api.server.Post("/login", api.Login)
	api.server.Get("/auth", api.Auth)

	api.address = fmt.Sprintf("%s:%d", ip, port)

	return &api
}

func (api *Server) Run() {
	logrus.Fatal(api.server.Listen(api.address))
}
