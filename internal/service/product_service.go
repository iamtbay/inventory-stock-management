package service

import (
	"context"

	"github.com/iamtbay/is-management/internal/domain"
	"github.com/iamtbay/is-management/pkg/helpers"
)

type ProductService struct {
	productRepository domain.ProductRepository
}

func NewProductService(productRepository domain.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

// PRODUCTS
func (p *ProductService) CreateProduct(product *domain.Product, ctx context.Context) error {
	product.ID = helpers.GenerateUUID()
	return p.productRepository.Save(product, ctx)
}

func (p *ProductService) FindProductByID(id string, ctx context.Context) (*domain.Product, error) {
	return p.productRepository.FindByID(id, ctx)
}

func (p *ProductService) UpdateStock(id string, stockQuantity int, ctx context.Context) (*domain.Product, error) {
	return p.productRepository.UpdateStock(id, stockQuantity, ctx)
}

func (p *ProductService) FindAll(ctx context.Context) ([]domain.Product, error) {
	return p.productRepository.FindAll(ctx)
}
