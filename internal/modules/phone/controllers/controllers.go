package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"phonebook/internal/modules/phone/requests/phone"
	"phonebook/internal/modules/phone/responses"
	"phonebook/internal/modules/phone/services"
	"strconv"
)

type Controller struct {
	phoneService services.PhoneServiceInterface
}

func New() *Controller {
	return &Controller{
		phoneService: services.New(),
	}
}

func (controller *Controller) GetPhone(c echo.Context) error {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	responsePhone := controller.phoneService.GetPhone(i)
	return c.JSON(http.StatusOK, responsePhone)
}

func (controller *Controller) GetPhones(userId int) responses.Phones {
	phones := controller.phoneService.GetPhones(userId)
	return phones
}

func (controller *Controller) GetTopTenPhones(userId int) responses.Phones {
	phones := controller.phoneService.GetTopTenPhones(userId)
	return phones
}

func (controller *Controller) CreatePhone(phoneRequest phone.StoreRequest, userId int) responses.Phone {
	newPhone := controller.phoneService.CreatePhone(phoneRequest, userId)
	return newPhone
}
