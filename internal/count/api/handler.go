package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (srv *Server) GetCounter(e echo.Context) error {
	msg, err := srv.uc.SelectCounter()
	if err != nil {
		return e.String(http.StatusInternalServerError, msg)
	}
	return e.JSON(http.StatusOK, msg)
}

func (srv *Server) SetCounter(e echo.Context) error {
	num := e.QueryParam("num")
	msg, err := srv.uc.SetCounter(num)
	if err != nil {
		return e.String(http.StatusBadRequest, msg)
	}
	return e.JSON(http.StatusOK, msg)
}

func (srv *Server) PostCounter(e echo.Context) error {
	msg, err := srv.uc.PostCounter()
	if err != nil {
		return e.String(http.StatusInternalServerError, msg)
	}
	return e.JSON(http.StatusOK, msg)
}

func (srv *Server) ClearCounter(e echo.Context) error {
	msg, err := srv.uc.ClearCounter()
	if err != nil {
		return e.String(http.StatusInternalServerError, msg)
	}
	return e.JSON(http.StatusOK, msg)
}
