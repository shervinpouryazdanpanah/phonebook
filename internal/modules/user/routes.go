package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Routes(e *echo.Echo) {
	g := e.Group("/user")
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "user get")
	})
}
