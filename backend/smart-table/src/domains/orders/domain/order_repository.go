package domain

import (
	"context"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
	"github.com/google/uuid"
)

type OrderRepository interface {
	Save(ctx context.Context, order utils.SharedRef[Order]) error
	Begin(ctx context.Context) error
	Commit(ctx context.Context) error

	FindOrders(ctx context.Context, orderUuids []uuid.UUID) ([]utils.SharedRef[Order], error)
	FindOrder(ctx context.Context, orderUuid uuid.UUID) (utils.SharedRef[Order], error)
	FindActiveOrderByTableId(ctx context.Context, tableId string) (utils.SharedRef[Order], error)
}
