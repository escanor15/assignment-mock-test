package service

import (
	"go_bootcamp/H8-mock/entity"
	"go_bootcamp/H8-mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetAllProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindAll").Return(nil)

	product, err := productService.GetAllProduct()

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetAllProduct(t *testing.T) {
	products := []*entity.Product{
		{Id: "1", Name: "Kaca mata"},
		{Id: "2", Name: "Laptop"},
	}

	productRepository.Mock.On("FindAll").Return(products)
	result, err := productService.GetAllProduct()

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, products, result, "result has to be products")
}
