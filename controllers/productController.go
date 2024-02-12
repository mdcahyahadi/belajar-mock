package controllers

import (
	"errors"
	"httpauth/interfaces"
	"httpauth/models"
)

type ProductService struct {
	Repository interfaces.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*models.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")

	}
	return product, nil
}

func (service ProductService) GetAllProduct() ([]*models.Product, error) {
	products := service.Repository.FindAll()
	if len(products) == 0 {
		return nil, errors.New("product not found")
	}
	return products, nil
}
