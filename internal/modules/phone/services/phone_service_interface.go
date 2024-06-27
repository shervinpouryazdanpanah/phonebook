package services

import (
	"phonebook/internal/modules/phone/requests/phone"
	"phonebook/internal/modules/phone/responses"
)

type PhoneServiceInterface interface {
	GetTopTenPhones(userId int) responses.Phones
	GetPhones(userId int) responses.Phones
	GetPhone(id int) responses.Phone
	CreatePhone(phoneRequest phone.StoreRequest, userId int) responses.Phone
}
