package service

import "customerService/app/models"

func GetList() []models.Customer {
	list := []models.Customer{
		models.Customer{
			ID:   1,
			Name: "Customer 1",
			City: "Samsun",
		},
		models.Customer{
			ID:   2,
			Name: "Customer 2",
			City: "İstanbul",
		},
		models.Customer{
			ID:   3,
			Name: "Customer 3",
			City: "Ankara",
		},
	}
	return list
}
