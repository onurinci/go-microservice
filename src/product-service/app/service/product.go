package service

import (
	"productService/app/models"
)

func GetList() []models.Product {
	list := []models.Product{
		models.Product{
			ID:   1,
			Code: "A-123",
			Name: "Product 1",
		},
		models.Product{
			ID:   2,
			Code: "B-123",
			Name: "Product 2",
		},
		models.Product{
			ID:   3,
			Code: "B-123",
			Name: "Product 3",
		},
	}
	return list
}
