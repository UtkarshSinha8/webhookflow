package repository

import (
	"context"

	"WebhookFlow/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(
	ctx context.Context,
	order *models.Order,
) error {

	query := `
	INSERT INTO orders (
		id,
		merchant_id,
		amount,
		status
	)
	VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		order.ID,
		order.MerchantID,
		order.Amount,
		order.Status,
	)

	return err
}
