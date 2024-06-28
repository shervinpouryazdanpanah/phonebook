package routes

import (
	"github.com/labstack/echo/v4"
	"phonebook/internal/modules/phone/controllers"
)

func Routes(e *echo.Echo) {
	phoneController := controllers.New()
	g := e.Group("/phones")
	g.GET("/:id", func(c echo.Context) error {
		return phoneController.GetPhone(c)
	})
}
