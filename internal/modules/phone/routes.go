package phone

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Routes(e *echo.Echo) {

	e.Group("/user")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "user get")
	})
}
