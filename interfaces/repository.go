package interfaces

import "httpauth/models"

type ProductRepository interface {
	FindById(id string) *models.Product
	FindAll() []*models.Product
}
