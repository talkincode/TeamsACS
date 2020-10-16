package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *HttpHandler) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
