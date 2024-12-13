package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (srv *Server) GetQuery(e echo.Context) error {
	msg, err := srv.uc.GetQuery()
	if err != nil {
		return e.String(http.StatusInternalServerError, msg)
	}
	return e.JSON(http.StatusOK, msg)
}

func (srv *Server) PostQuery(e echo.Context) error {
	name := e.QueryParam("name")
	age := e.QueryParam("age")
	msg, err := srv.uc.PostQuery(name, age)
	if err != nil {
		return e.String(http.StatusInternalServerError, msg)
	}
	return e.JSON(http.StatusCreated, msg)
}

func (srv *Server) ClearQuery(e echo.Context) error {
	msg, err := srv.uc.ClearQuery()
	if err != nil {
		return e.String(http.StatusInternalServerError, msg)
	}
	return e.JSON(http.StatusOK, msg)
}
