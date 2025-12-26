package service

import (
	"context"
	"errors"

	"github.com/iamtbay/is-management/internal/domain"
	"github.com/iamtbay/is-management/pkg/helpers"
)

type OrderService struct {
	orderRepository   domain.OrderRepository
	productRepository domain.ProductRepository
}

func NewOrderService(orderRepository domain.OrderRepository, productRepository domain.ProductRepository) *OrderService {
	return &OrderService{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (s *OrderService) CreateOrder(order *domain.Order, ctx context.Context) error {
	checkStock, err := s.productRepository.FindByID(order.ProductID, ctx)
	if err != nil {
		return err
	}
	if checkStock.Stock < order.Quantity {
		return errors.New("stock is not enough")
	}
	_, err = s.productRepository.UpdateStock(order.ProductID, order.Quantity, ctx)
	if err != nil {
		return err
	}
	order.TotalPrice = checkStock.Price * float64(order.Quantity)
	order.ID = helpers.GenerateUUID()
	return s.orderRepository.Save(order, ctx)
}

func (s *OrderService) FindAll(ctx context.Context) ([]domain.Order, error) {
	return s.orderRepository.FindAll(ctx)
}

func (s *OrderService) FindByID(id string, ctx context.Context) (*domain.Order, error) {
	return s.orderRepository.FindByID(id, ctx)
}
