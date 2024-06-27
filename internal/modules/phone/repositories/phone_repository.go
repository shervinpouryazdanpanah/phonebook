package repositories

import (
	"gorm.io/gorm"
	"phonebook/internal/modules/phone/models"
	"phonebook/pkg/database"
)

type PhoneRepository struct {
	DB *gorm.DB
}

func New() *PhoneRepository {
	return &PhoneRepository{
		DB: database.Connection(),
	}
}

func (phoneRepository *PhoneRepository) List(limit int, userId int) []models.Phone {
	var phoneList []models.Phone
	phoneRepository.DB.Limit(limit).Joins("User").Order("created_at desc").
		Find(&phoneList, "user_id = ?", userId)
	return phoneList
}

func (phoneRepository *PhoneRepository) All(userId int) []models.Phone {
	var phoneList []models.Phone
	phoneRepository.DB.Joins("User").Order("created_at desc").
		Find(&phoneList, "user_id = ?", userId)
	return phoneList
}

func (phoneRepository *PhoneRepository) Find(id int) models.Phone {
	var phone models.Phone
	phoneRepository.DB.Joins("User").First(&phone, id)
	return phone
}

func (phoneRepository *PhoneRepository) Create(phone models.Phone) models.Phone {
	var newPhone models.Phone
	result := phoneRepository.DB.Create(phone)
	if result.Error != nil {
		result.Scan(&newPhone)
	}
	return newPhone
}
