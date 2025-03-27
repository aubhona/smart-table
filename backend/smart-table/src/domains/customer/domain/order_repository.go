package domain

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/smart-table/src/utils"
)

type OrderRepository interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Commit(ctx context.Context, tx pgx.Tx) error

	Save(ctx context.Context, tx pgx.Tx, order utils.SharedRef[Order]) error

	FindOrders(ctx context.Context, orderUUIDs []uuid.UUID) ([]utils.SharedRef[Order], error)
	FindOrder(ctx context.Context, orderUUID uuid.UUID) (utils.SharedRef[Order], error)
	FindActiveOrderByTableID(ctx context.Context, tableID string) (utils.SharedRef[Order], error)
}
