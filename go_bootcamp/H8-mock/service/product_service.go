package service

import (
	"errors"
	"go_bootcamp/H8-mock/entity"
	"go_bootcamp/H8-mock/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (service ProductService) GetAllProduct() ([]*entity.Product, error) {
	products := service.Repository.FindAll()
	if len(products) == 0 {
		return nil, errors.New("No product found")
	}
	return products, nil
}
