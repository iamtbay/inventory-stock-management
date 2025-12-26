package service

import (
	"context"
	"testing"

	"github.com/iamtbay/is-management/internal/domain"
)

type mockProductRepo struct {
	fakeProduct *domain.Product
	fakeError   error
}

func (m *mockProductRepo) FindByID(id string, ctx context.Context) (*domain.Product, error) {
	return m.fakeProduct, m.fakeError
}

func (m *mockProductRepo) UpdateStock(id string, stockQuantity int, ctx context.Context) (*domain.Product, error) {
	if m.fakeProduct != nil {
		m.fakeProduct.Stock -= stockQuantity
	}
	return m.fakeProduct, m.fakeError
}

func (m *mockProductRepo) FindAll(ctx context.Context) ([]domain.Product, error) {
	return nil, nil
}

func (m *mockProductRepo) Save(product *domain.Product, ctx context.Context) error {
	return nil
}

// TESTS
func TestFindByID(t *testing.T) {
	mockPRepo := &mockProductRepo{fakeProduct: &domain.Product{ID: "prod-1", Name: "Laptop", Price: 100.0, Stock: 10}}
	svc := NewProductService(mockPRepo)
	product, err := svc.FindProductByID("prod-1", context.Background())
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if product.ID != "prod-1" {
		t.Errorf("expected product ID 'prod-1', got %v", product.ID)
	}
}

func TestUpdateStock(t *testing.T) {
	mockPRepo := &mockProductRepo{fakeProduct: &domain.Product{ID: "prod-1", Name: "Laptop", Price: 100.0, Stock: 10}}
	svc := NewProductService(mockPRepo)
	product, err := svc.UpdateStock("prod-1", 2, context.Background())
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if product.Stock != 8 {
		t.Errorf("expected stock 8, got %v", product.Stock)
	}
}

