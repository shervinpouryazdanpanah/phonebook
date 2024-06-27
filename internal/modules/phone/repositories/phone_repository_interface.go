package repositories

import "phonebook/internal/modules/phone/models"

type PhoneRepositoryInterface interface {
	List(limit int, userId int) []models.Phone
	All(userId int) []models.Phone
	Find(id int) models.Phone
	Create(phone models.Phone) models.Phone
}
