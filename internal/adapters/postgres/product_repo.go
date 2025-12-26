package postgres

import (
	"context"
	"errors"

	"github.com/iamtbay/is-management/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	conn *pgxpool.Pool
}

// NEW PRODUCT REPO
func NewProductRepository(conn *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{conn: conn}
}

// SAVE
func (r *ProductRepository) Save(product *domain.Product, ctx context.Context) error {
	query := `INSERT INTO products (id, name, price, stock) VALUES ($1, $2, $3, $4)`
	_, err := r.conn.Exec(ctx, query, product.ID, product.Name, product.Price, product.Stock)
	if err != nil {
		return err
	}
	return nil
}

// ! MUST DO
// FIND ALL
func (r *ProductRepository) FindAll(ctx context.Context) ([]domain.Product, error) {
	var query = `SELECT * FROM products`
	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// FIND BY ID
func (r *ProductRepository) FindByID(id string, ctx context.Context) (*domain.Product, error) {
	query := `SELECT * FROM products WHERE id=$1`
	var product domain.Product
	err := r.conn.QueryRow(ctx, query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

// UPDATE STOCK
func (r *ProductRepository) UpdateStock(id string, stockQuantity int, ctx context.Context) (*domain.Product, error) {
	query := `UPDATE products SET stock=stock-$2 WHERE id=$1 AND stock>=$2 RETURNING *`
	var product domain.Product
	err := r.conn.QueryRow(ctx, query, id, stockQuantity).Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("stock is not enough")
		}
		return nil, err
	}
	return &product, nil
}
