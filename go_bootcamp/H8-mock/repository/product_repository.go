package repository

import "go_bootcamp/H8-mock/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
	FindAll() []*entity.Product
}
