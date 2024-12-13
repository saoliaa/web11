package api

import (
	"fmt"

	"github.com/ValeryBMSTU/web-11/internal/auth/middleware"
	"github.com/labstack/echo/v4"
)

type Server struct {
	maxSize int

	server  *echo.Echo
	address string

	uc Usecase
}

func NewServer(ip string, port int, maxSize int, uc Usecase) *Server {
	api := Server{
		maxSize: maxSize,
		uc:      uc,
	}

	api.server = echo.New()
	api.server.Use(middleware.AuthMiddleware)

	api.server.GET("/count", api.GetCounter)
	api.server.POST("/count", api.PostCounter)
	api.server.PUT("/count", api.SetCounter)
	api.server.DELETE("/count", api.ClearCounter)

	api.address = fmt.Sprintf("%s:%d", ip, port)

	return &api
}

func (api *Server) Run() {
	api.server.Logger.Fatal(api.server.Start(api.address))
}
