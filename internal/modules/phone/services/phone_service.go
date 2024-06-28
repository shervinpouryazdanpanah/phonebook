package services

import (
	phoneModel "phonebook/internal/modules/phone/models"
	PhoneRepository "phonebook/internal/modules/phone/repositories"
	"phonebook/internal/modules/phone/requests/phone"
	"phonebook/internal/modules/phone/responses"
)

type PhoneService struct {
	phoneRepository PhoneRepository.PhoneRepositoryInterface
}

func New() *PhoneService {
	return &PhoneService{
		phoneRepository: PhoneRepository.New(),
	}
}

func (phoneService *PhoneService) GetTopTenPhones(userId int) responses.Phones {
	phones := phoneService.phoneRepository.List(10, userId)
	return responses.ToPhones(phones)
}

func (phoneService *PhoneService) GetPhones(userId int) responses.Phones {
	phones := phoneService.phoneRepository.All(userId)
	return responses.ToPhones(phones)
}

func (phoneService *PhoneService) GetPhone(id int) responses.Phone {
	phone := phoneService.phoneRepository.Find(id)
	return responses.ToPhone(phone)
}

func (phoneService *PhoneService) CreatePhone(phoneRequest phone.StoreRequest, userId int) responses.Phone {
	var model phoneModel.Phone
	model.UserId = userId
	model.Number = phoneRequest.Number
	newPhone := phoneService.phoneRepository.Create(model)
	return responses.ToPhone(newPhone)
}
