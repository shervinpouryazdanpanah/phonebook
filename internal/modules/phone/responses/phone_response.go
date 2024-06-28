package responses

import (
	phoneModel "phonebook/internal/modules/phone/models"
	UserResponse "phonebook/internal/modules/user/responses"
)

type Phone struct {
	Number string            `json:"name" xml:"name"`
	User   UserResponse.User `json:"user" xml:"user"`
}

type Phones struct {
	Data []Phone `json:"data"`
}

func ToPhone(phone phoneModel.Phone) Phone {
	return Phone{
		Number: phone.Number,
		User:   UserResponse.ToUser(phone.User),
	}
}

func ToPhones(phones []phoneModel.Phone) Phones {
	var phonesData Phones
	for _, phone := range phones {
		phonesData.Data = append(phonesData.Data, ToPhone(phone))
	}
	return phonesData
}
