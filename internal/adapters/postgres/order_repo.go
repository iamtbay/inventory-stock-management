package postgres

import (
	"context"

	"github.com/iamtbay/is-management/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	conn *pgxpool.Pool
}

// NEW ORDER REPO
func NewOrderRepository(conn *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{conn: conn}
}

func (r *OrderRepository) Save(order *domain.Order, ctx context.Context) error {
	var query = `INSERT INTO orders (id, product_id, quantity,total_price) VALUES ($1, $2, $3,$4)`

	_, err := r.conn.Exec(ctx, query, order.ID, order.ProductID, order.Quantity, order.TotalPrice)
	if err != nil {
		return err
	}

	return nil
}

// ORDERS MUST DO
func (r *OrderRepository) FindAll(ctx context.Context) ([]domain.Order, error) {
	var query = `SELECT * FROM orders`
	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.ID, &order.ProductID, &order.Quantity, &order.TotalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepository) FindByID(id string, ctx context.Context) (*domain.Order, error) {
	var query = `SELECT * FROM orders WHERE id = $1`
	row := r.conn.QueryRow(ctx, query, id)
	var order domain.Order
	if err := row.Scan(&order.ID, &order.ProductID, &order.Quantity, &order.TotalPrice); err != nil {
		return nil, err
	}
	return &order, nil
}
