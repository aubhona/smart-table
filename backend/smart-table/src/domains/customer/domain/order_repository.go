package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type OrderRepository interface {
	Save(ctx context.Context, order utils.SharedRef[Order]) error
	Begin(ctx context.Context) error
	Commit(ctx context.Context) error

	FindOrders(ctx context.Context, orderUUIDs []uuid.UUID) ([]utils.SharedRef[Order], error)
	FindOrder(ctx context.Context, orderUUID uuid.UUID) (utils.SharedRef[Order], error)
	FindActiveOrderByTableID(ctx context.Context, tableID string) (utils.SharedRef[Order], error)
}
