package repository

import (
	"context"
	"database/sql"

	"github.com/Hamiduzzaman96/Product---Service/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetByID(ctx context.Context, id int64) (*domain.Product, error) {
	query := `SELECT id,name,description,price,stock,created_at FROM products WHERE id =?`

	row := r.db.QueryRowContext(ctx, query, id)

	var p domain.Product

	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Stock,
		&p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &p, nil
}
