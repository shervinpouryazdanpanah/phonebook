package repositories

import "phonebook/internal/modules/phone/models"

type PhoneRepositoryInterface interface {
	List(limit int, userId int) []models.Phone
	find(id int) models.Phone
	create(phone models.Phone) models.Phone
}
