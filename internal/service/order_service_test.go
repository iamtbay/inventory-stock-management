package service

import (
	"context"
	"testing"

	"github.com/iamtbay/is-management/internal/domain"
)

type mockOrderRepo struct {
	saveCalled bool
}

func (m *mockOrderRepo) Save(order *domain.Order, ctx context.Context) error {
	m.saveCalled = true
	return nil
}

func (m *mockOrderRepo) FindAll(ctx context.Context) ([]domain.Order, error) {
	return nil, nil
}

func (m *mockOrderRepo) FindByID(id string, ctx context.Context) (*domain.Order, error) {
	return nil, nil
}

// TESTS
func TestCreateOrder_Success(t *testing.T) {
	existingProduct := &domain.Product{
		ID:    "prod-1",
		Name:  "Laptop",
		Price: 100.0,
		Stock: 10,
	}

	mockPRepo := &mockProductRepo{
		fakeProduct: existingProduct,
	}
	mockORRepo := &mockOrderRepo{}

	svc := NewOrderService(mockORRepo, mockPRepo)

	order := &domain.Order{
		ProductID: "prod-1",
		Quantity:  2,
	}
	err := svc.CreateOrder(order, context.Background())
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if order.TotalPrice != 200.0 {
		t.Errorf("expected total price 200.0, got %v", order.TotalPrice)
	}
	if !mockORRepo.saveCalled {
		t.Errorf("Order repository Save method was not called")
	}
	if existingProduct.Stock != 8 {
		t.Errorf("expected stock 8, got %v", existingProduct.Stock)
	}
}

func TestCreateOrder_InsufficientStock(t *testing.T) {
	existingProduct := &domain.Product{
		ID:    "prod-1",
		Stock: 1,
	}

	mockPRepo := &mockProductRepo{
		fakeProduct: existingProduct,
	}
	mockORRepo := &mockOrderRepo{}
	svc := NewOrderService(mockORRepo, mockPRepo)

	order := &domain.Order{
		ProductID: "prod-1",
		Quantity:  5,
	}
	err := svc.CreateOrder(order, context.Background())
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if err.Error() != "stock is not enough" {
		t.Errorf("expected error message 'stock is not enough', got %v", err.Error())
	}
}
